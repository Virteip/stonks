package database

import (
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/cockroachdb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations runs database migrations
func RunMigrations(connectionString string) error {
	cockroachConnectionString := strings.Replace(connectionString, "postgresql://", "cockroachdb://", 1)

	m, err := migrate.New(
		"file://cmd/migrations",
		cockroachConnectionString,
	)
	if err != nil {
		return fmt.Errorf("migration error: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
