package database

import (
	"fmt"

	"onlinevoting/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=piyush@6395 dbname=ecommerce port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db

	db.AutoMigrate(
		&models.User{},
		&models.Election{},
		&models.Candidate{},
		&models.Vote{},
	)

	fmt.Println("Database connected successfully!")
}
