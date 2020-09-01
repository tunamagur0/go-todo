package db

import (
	"log"
	"os"

	"github.com/tunamagur0/go-todo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
