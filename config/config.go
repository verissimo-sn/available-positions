package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitSQLite()
	if err != nil {
		return fmt.Errorf("Error initializing database: %v", err)
	}
	return nil
}

func GetLogger(prefix string) *Logger {
	return NewLogger(prefix)
}

func GetSQLite() *gorm.DB {
	return db
}
