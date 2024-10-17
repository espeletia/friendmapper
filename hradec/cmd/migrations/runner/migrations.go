package runner

import (
	"context"
	"hradec/internal/config"
	"hradec/internal/setup"

	"github.com/pressly/goose/v3"
)

func RunMigrations() error {
	configuration := config.LoadConfig()
	return RunMigrationsWithConfiguration(*configuration)
}

func RunMigrationsWithConfiguration(configuration config.Config) error {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "config", configuration)

	dBconn, err := setup.SetupDb(&configuration)
	if err != nil {
		return err
	}

	if err := goose.Up(dBconn, configuration.MigrationsConfig.MigrationPath); err != nil {
		return err
	}

	return nil
}
