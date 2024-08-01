package database

import (
	"fmt"
	"someName/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=UTC"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	
	DB = connection

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		panic(fmt.Sprintf("Failed to auto migrate: %v", err))
	}

	fmt.Println("Successfully connected to the database")
}