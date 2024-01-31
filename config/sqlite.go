package config

import (
	"os"

	"github.com/verissimo-sn/available-positions/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	filePath := "./db/sqlite.db"
	err := createDatabaseFile(logger, filePath)
	if err != nil {
		return nil, err
	}
	dbInstance, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Sqlite opening error: %v", err)
		return nil, err
	}
	err = dbInstance.AutoMigrate(&schemas.AvailablePosition{})
	if err != nil {
		logger.Errorf("Sqlite auto migration error error: %v", err)
		return nil, err
	}
	return dbInstance, nil
}

func createDatabaseFile(logger *Logger, filePath string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		logger.Info("Sqlite database file not found, creating a new one")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return err
		}
		file, err := os.Create(filePath)
		if err != nil {
			logger.Errorf("Sqlite database file creation error: %v", err)
			return err
		}
		file.Close()
	}
	return nil
}
