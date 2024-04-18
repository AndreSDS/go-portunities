package config

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"

	database "github.com/andresds/go-portunities/internal/db"
)

func InitializeSQLite() (*database.Queries, error) {
	// initialize logger
	logger := GetLogger("sqlite")

	const dbPath = "./sql/opennings.db"

	// check if database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")
		// create database file
		err = os.MkdirAll("./sql", os.ModePerm)
		if err != nil {
			logger.Errorf("Error creating database directory: %v", err.Error())
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Errorf("Error creating database file: %v", err.Error())
			return nil, err
		}
		logger.Info("Database file created")
		file.Close()
	}

	// connect to database
	dbConn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err.Error())
		return nil, err
	}
	defer dbConn.Close()

	db := database.New(dbConn)

	return db, nil
}
