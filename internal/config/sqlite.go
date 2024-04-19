package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/andresds/go-portunities/schemas"
)

func InitializeSQLite() (*gorm.DB, error) {
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
	db, err := gorm.Open(sqlite.Open(dbPath))
	if err != nil {
		logger.Errorf("Error connecting to database: %v", err.Error())
		return nil, err
	}

	// auto migrate the schema
	err = db.AutoMigrate(&schemas.Openning{})
	if err != nil {
		logger.Errorf("Error migrating schema: %v", err.Error())
		return nil, err
	}

	logger.Info("Connected to database")

	return db, nil
}
