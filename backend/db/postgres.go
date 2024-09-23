package db

import (
	"decentralised_payment_gateway/models" // Import the models package
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB //Declares a global variable DB of type *gorm.DB

func ConnectPostgres() { //This function establishes a connection to the PostgreSQL database.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	var err error //Attempts to open a connection to the database using GORM. If it fails, it panics and prints an error message.
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	fmt.Println("Database connected!")
}

func AutoMigrate() { //takes care of creating or updating the tables for Merchant and Payment based on the defined struct fields in the models package.
	DB.AutoMigrate(&models.Merchant{}, &models.Payment{}) // Using models here
}
