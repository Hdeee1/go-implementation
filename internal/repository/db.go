package repository

import (
	"database/sql"

	"github.com/Hdeee1/go-implementation/config"
)

var DB *sql.DB

func ConnectDB(cfg *config.Config) error {
	dsn := cfg.GetDSN()

	db, err := sql.Open("postgres", dsn)
	err = db.Ping()
	DB = db
	return  err
}