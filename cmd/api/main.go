package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/betology/selvaGo/internal/config"
	"github.com/betology/selvaGo/internal/database"
	"github.com/betology/selvaGo/internal/handlers"
	"github.com/betology/selvaGo/internal/routes"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("./internal/config")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize the database connection.
	db, err := database.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize Gin router.
	router := gin.Default()

	// Create a new instance of the handlers
	nombreHandler := handlers.NewNombreHandler(db)

	// Setup routes
	routes.SetupNombreRoutes(router, nombreHandler)

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = cfg.Server.Port
	}
	router.Run(":" + port)
}
