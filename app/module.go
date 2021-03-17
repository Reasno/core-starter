package app

import (
	"github.com/DoNewsCode/core-starter/app/commands"
	"github.com/DoNewsCode/core-starter/app/databases"
	"github.com/DoNewsCode/core/contract"
	"github.com/DoNewsCode/core/di"
	"github.com/DoNewsCode/core/otgorm"
	"github.com/spf13/cobra"
)

func New(config contract.ConfigAccessor) Module {
	return Module{config: config}
}

func Providers() di.Deps {
	return []interface{}{
		func(maker otgorm.Maker) *DB {
			db, _ := maker.Make("default")
			return (*DB)(db)
		},
		func(maker otgorm.Maker) *FooDB {
			db, _ := maker.Make("foo")
			return (*FooDB)(db)
		},
		func(maker otgorm.Maker) *BarDB {
			db, _ := maker.Make("bar")
			return (*BarDB)(db)
		},
	}
}

type Module struct {
	config contract.ConfigAccessor
}

func (m Module) ProvideSeed() []*otgorm.Seed {
	return databases.Seeders()
}

func (m Module) ProvideMigration() []*otgorm.Migration {
	return databases.Migrations()
}

func (m Module) ProvideCommand(command *cobra.Command) {
	command.AddCommand(
		commands.NewVersionCommand(m.config),
	)
}
