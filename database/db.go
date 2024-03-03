package database

import (
	"log"

	"github.com/flpcastro/apirest-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DatabaseConnect() {
	connString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connString))
	if err != nil {
		log.Panic("Error connecting to database: ", err)
	}

	DB.AutoMigrate(&models.Student{})
}
