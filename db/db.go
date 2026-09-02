package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var err error

type Movie struct {
	Id string `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
	Description string `json:"description"`
}

func InitPostgresDB() {
	err = godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error loading database: %s", err);
		return
	}

	var (
		host = os.Getenv("DB_HOST")
		port = os.Getenv("DB_PORT")
		dbUser = os.Getenv("DB_USER")
		dbName = os.Getenv("DB_NAME")
		password = os.Getenv("DB_PASSWORD")
	)

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		dbUser,
		dbName,
		password,
	)

	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		fmt.Print("Error opening connection to db: s%", err)
		return
	}

	DB.Exec("SELECT 'CREATE DATABASE ?' WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname=?)", dbName, dbName)
	DB.AutoMigrate(Movie{})
}