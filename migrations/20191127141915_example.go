package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upExample, downExample)
}

func upExample(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`
	create table example
	(
		created_at 		 timestamp default now()
	);
	`)
	return err
}

func downExample(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("drop table example;")
	return err
}
