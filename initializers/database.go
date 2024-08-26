package initializers

import (
	"log"
	"os"

	"github.com/rashidkalwar/go-crud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}

	// Auto-migrate models
	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}
}
