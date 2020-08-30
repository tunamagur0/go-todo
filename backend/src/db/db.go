package db

import (
	"log"

	"github.com/tunamagur0/go-todo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(dbPath string) *gorm.DB {
	var db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
