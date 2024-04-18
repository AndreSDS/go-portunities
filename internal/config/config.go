package config

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db     *sql.DB
	logger *Logger
)

func Init() error {
	return errors.New("not implemented")
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
