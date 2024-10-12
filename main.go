package main

import (
	"go-boilerplate/models"
	"go-boilerplate/routes"
	"go-boilerplate/services"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to Database
	dsn := os.Getenv("DB_DSN")
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// AutoMigrate User model
	db.AutoMigrate(&models.User{})

	// Set database reference in services
	services.SetDB(db)

	// Initialize Gin Router
	r := gin.Default()

	// Set up routes
	routes.SetupRoutes(r, db)

	gin.SetMode(os.Getenv("GIN_MODE"))

	// Start the server
	r.Run() // default listens on :8080
}
