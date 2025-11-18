package migration

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
)

type Migrator struct {
	db *sql.DB
}

func NewMigrator(db *sql.DB) *Migrator {
	return &Migrator{db: db}
}

// Run выполняет все непримененные миграции
func (m *Migrator) Run() error {
	goose.SetBaseFS(nil)

	ctx := context.Background()

	if err := goose.UpContext(ctx, m.db, "migrations"); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	version, err := goose.GetDBVersionContext(ctx, m.db)
	if err != nil {
		return fmt.Errorf("failed to get DB version: %w", err)
	}

	log.Printf("Migrations applied successfully. Current version: %d", version)
	return nil
}

// Status проверяет статус миграции - какие версии приминены, а какие - нет.
func (m *Migrator) Status() error {
	ctx := context.Background()
	return goose.StatusContext(ctx, m.db, "migrations")
}
