package app

import (
	"github.com/DoNewsCode/core-starter/app/commands"
	pb "github.com/DoNewsCode/core-starter/app/proto"
	"github.com/DoNewsCode/core-starter/app/service"
	"github.com/DoNewsCode/core/contract"
	"github.com/DoNewsCode/core/di"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func New(config contract.ConfigAccessor, app pb.AppServer) Module {
	return Module{config: config, app: app}
}

func Providers() di.Deps {
	return []interface{}{
		service.NewAppService,
	}
}

type Module struct {
	config contract.ConfigAccessor
	app    pb.AppServer
}

func (m Module) ProvideGRPC(server *grpc.Server) {
	// @TODO apply middlewares
	pb.RegisterAppServer(server, m.app)
}

func (m Module) ProvideHTTP(router *mux.Router) {
	router.PathPrefix("/").Handler(pb.NewAppHandler(m.app, http.Middleware(middleware.Chain(
		recovery.Recovery(),
		logging.Server(),
	))))
}

func (m Module) ProvideCommand(command *cobra.Command) {
	command.AddCommand(
		commands.NewVersionCommand(m.config),
	)
}
