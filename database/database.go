package database

import (
	"log"

	"gitlab.com/alura-courses-code/golang/go-gin-rest-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Fatalln("Failed to connect to database:", err)
	}

	DB.AutoMigrate(&models.Student{})
}
