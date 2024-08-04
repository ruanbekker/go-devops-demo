package config

import (
	"fmt"
	"os"

	"github.com/go-devops-demo/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	var dsn string

	switch os.Getenv("DB_TYPE") {
	case "postgresql":
		dsn = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		dbStorage := os.Getenv("DB_STORAGE")
		if dbStorage == "memory" {
			dsn = "file::memory:?cache=shared"
		} else {
			dsn = dbStorage
		}
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	default:
		panic("Unsupported DB_TYPE. Supported types are 'postgresql' and 'sqlite'")
	}

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB.AutoMigrate(&models.User{})
}
