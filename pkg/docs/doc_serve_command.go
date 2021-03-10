package docs

import (
    "context"
    "github.com/DoNewsCode/core/container"
    "github.com/DoNewsCode/core/contract"
    "github.com/DoNewsCode/core/di"
    "github.com/DoNewsCode/core/logging"
    "github.com/go-kit/kit/log"
    "github.com/gorilla/mux"
    "github.com/oklog/run"
    "github.com/pkg/errors"
    "github.com/spf13/cobra"
    "net"
    "net/http"
    "os"
    "os/signal"
    "syscall"
)

type docServeIn struct {
    di.In

    Config     contract.ConfigAccessor
    Logger     log.Logger
    Container  contract.Container
    HttpServer *http.Server `optional:"true"`
}

func NewDocServeModule(in docServeIn) docServeModule {
    return docServeModule{
        in,
    }
}

var _ container.CommandProvider = (*docServeModule)(nil)

type docServeModule struct {
    in docServeIn
}

func (s docServeModule) ProvideCommand(command *cobra.Command) {
    command.AddCommand(newServeCmd(s.in))
}

func newServeCmd(p docServeIn) *cobra.Command {
    var serveCmd = &cobra.Command{
        Use:   "docs",
        Short: "Start the docs server",
        Long:  `Start the docs server.`,
        RunE: func(cmd *cobra.Command, args []string) error {
            doc := ""

            if len(args) > 0 {
                doc = "/" + args[0]
            }

            var (
                g run.Group
                l = logging.WithLevel(p.Logger)
            )

            // Start HTTP server
            {
                httpAddr := p.Config.String("http.addr")
                ln, err := net.Listen("tcp", httpAddr)
                if err != nil {
                    return errors.Wrap(err, "failed start http server")
                }
                if p.HttpServer == nil {
                    p.HttpServer = &http.Server{}
                }
                router := mux.NewRouter()
                router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./docs"+doc))))
                p.HttpServer.Handler = router
                g.Add(func() error {
                    l.Infof("http service is listening at %s", ln.Addr())
                    return p.HttpServer.Serve(ln)
                }, func(err error) {
                    _ = p.HttpServer.Shutdown(context.Background())
                    _ = ln.Close()
                })
            }

            // Graceful shutdown
            {
                sig := make(chan os.Signal, 1)
                signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
                g.Add(func() error {
                    select {
                    case s := <-sig:
                        l.Errf("signal received: %s", s)
                    case <-cmd.Context().Done():
                        l.Errf(cmd.Context().Err().Error())
                    }
                    return nil
                }, func(err error) {
                    close(sig)
                })
            }

            // Additional run groups
            p.Container.ApplyRunGroup(&g)

            if err := g.Run(); err != nil {
                return err
            }

            l.Infof("graceful shutdown complete; see you next time :)")
            return nil
        },
    }
    return serveCmd
}
