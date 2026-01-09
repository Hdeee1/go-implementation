package migrations

import (
	"database/sql"
	"os"
)

func Migrate(db *sql.DB) error {
	query, err := os.ReadFile("./migrations/001_create_student_table.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(query))
	if err != nil {
		return err
	}
	return nil
}