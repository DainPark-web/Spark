package database

import (
	"fmt"
	"log"
	"os"
	"spark/config"
	"spark/models"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	mode := os.Getenv("MODE")

	if mode == "dev" {
		// Development : SQLite
		db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		DB = db
	} else {
		// Production : Postgres
		config := config.GetDBConfig()
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config.Host, config.User, config.Password, config.DBName, config.Port)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}
		DB = db

	}

	// Auto Migration
	err := DB.AutoMigrate(
		&models.User{},
		// &models.ChatRoom{},
		// &models.Message{},
		// &models.Like{},
		// &models.PartnerMatch{},
		// &models.FeedbackQuestion{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
