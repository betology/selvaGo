package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"go_selva/internal/api"
	"go_selva/internal/db"
)

func main() {
	// Initialize database connection
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize Gin router
	router := gin.Default()

	// Initialize API handlers
	apiHandler := api.NewAPIHandler(database)

	// Define API routes
	apiGroup := router.Group("/nombres")
	{
		apiGroup.GET("/search", apiHandler.SearchNombres)
		apiGroup.POST("", apiHandler.CreateNombre)
		apiGroup.GET("", apiHandler.GetNombres)
		apiGroup.GET("/:id", apiHandler.GetNombreByID)
		apiGroup.PUT("/:id", apiHandler.UpdateNombre)
		apiGroup.DELETE("/:id", apiHandler.DeleteNombre)
	}

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
