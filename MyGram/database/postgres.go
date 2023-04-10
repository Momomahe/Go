package database

import (
	"fmt"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "catwater888"
	DB_PORT     = 5432
	DB_NAME     = "postgres"
	DEBUG_MODE  = true // true/false
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.User{}, models.SocialMedia{}, models.Photo{}, models.Comment{})
}

func GetDB() *gorm.DB {
	if DEBUG_MODE {
		return db.Debug()
	}

	return db
}
