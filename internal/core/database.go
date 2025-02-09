package core

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"

	"github.com/prakoso-id/go-windsurf/internal/domain/models"
)

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Add health check
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			if err := sqlDB.Ping(); err != nil {
				log.Println("Database connection lost:", err)
			}
		}
	}()

	// Auto migrate models
	db.AutoMigrate(&models.Product{})

	return db
}
