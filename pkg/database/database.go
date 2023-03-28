package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"love-scroll-api/internal/config"
	"love-scroll-api/internal/model"
)

type DB = gorm.DB

func Connect(cfg *config.Config) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.Database.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	return (*DB)(db), nil
}
