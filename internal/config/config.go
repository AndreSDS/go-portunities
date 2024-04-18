package config

import (
	database "github.com/andresds/go-portunities/internal/db"
)

var (
	db     *database.Queries
	logger *Logger
)

func Init() error {
	var err error

	// initialize sqlite
	db, err = InitializeSQLite()
	if err != nil {
		return err
	}
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetSQLite() *database.Queries {
	return db
}
