package config

import (
	"os"

	"github.com/dev-araujo/go-job/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	const DB_PATH string = "./db/main.db"

	logger := GetLogger("sqlite")

	_, err := os.Stat(DB_PATH)
	if os.IsNotExist(err) {
		logger.Info("database file not found, creating...")

		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.Create(DB_PATH)

	if err != nil {
		return nil, err
	}

	file.Close()

	db, err := gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		logger.Errorf("sqlite opening error %v", err)

		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("sqlite automigration error %v", err)
		return nil, err
	}

	return db, nil
}
