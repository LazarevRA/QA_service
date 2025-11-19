package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upCreateAnswersTable, downCreateAnswersTable)
}

func upCreateAnswersTable(ctx context.Context, tx *sql.Tx) error {
	query := `
    CREATE TABLE IF NOT EXISTS answers (
        id SERIAL PRIMARY KEY,
        question_id INTEGER NOT NULL,
        user_id UUID NOT NULL,
        text TEXT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT fk_question
            FOREIGN KEY(question_id) 
            REFERENCES questions(id)
            ON DELETE CASCADE
    );`
	_, err := tx.ExecContext(ctx, query)
	return err
}

func downCreateAnswersTable(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "DROP TABLE IF EXISTS answers;")
	return err
}
