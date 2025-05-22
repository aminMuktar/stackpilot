package database

import (
	"fmt"
	"os"

	"github.com/aminMuktar/stackpilot/internal/logger"
	"github.com/aminMuktar/stackpilot/internal/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatal("Failed to connect to the database",
			zap.Error(err),
			zap.String("dsn", dsn),
		)
	}

	logger.Log.Info("Database connection established")
	DB = db
	if err := db.AutoMigrate(
		&models.User{},
		// &models.Tenant{},
		// &models.Role{},
		// &models.Permission{},
		// Add other models here
	); err != nil {
		logger.Log.Fatal("Auto migration failed", zap.Error(err))
	}

}
