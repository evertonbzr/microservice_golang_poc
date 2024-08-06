package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Database(uri string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(uri), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
