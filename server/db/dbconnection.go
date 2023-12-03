package db

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/silastgoes/client-server-api/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("bid.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(models.Bid{})
	return db, nil
}
