package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateQuestionsTable, downCreateQuestionsTable)
}

func upCreateQuestionsTable(ctx context.Context, tx *sql.Tx) error {
	query := `
    CREATE TABLE IF NOT EXISTS questions (
        id SERIAL PRIMARY KEY,
        text TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
    );`
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downCreateQuestionsTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE IF EXISTS questions;")
	return err
}
