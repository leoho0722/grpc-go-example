package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(sqlite.Open("database/db.sqlite"))
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Album{})
	if err != nil {
		panic("failed to migrate database")
	}

	migrator := db.Migrator()
	migrator.HasTable(&Album{})

	DB = db
}
