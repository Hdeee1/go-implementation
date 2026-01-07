package repository

import (
	"database/sql"

	"github.com/Hdeee1/go-implementation/config"
)


func InitDB(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.GetDSN()

	db, err := sql.Open("postgres", dsn)
	if err !=nil {
		return nil, err
	} 

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}