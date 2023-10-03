package migrate

import (
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed schema/*.sql
var postgresqlFS embed.FS

func UP(dsn string) error {
	drv, err := iofs.New(postgresqlFS, "schema")
	if err != nil {
		return fmt.Errorf("create driver: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", drv, dsn)
	if err != nil {
		return fmt.Errorf("create migration engine instance: %w", err)
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrating up: %w", err)
	}

	return nil
}
