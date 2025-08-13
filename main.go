package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"kalkulatorOkeFinansial/routes"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	// Middleware

	// rute
	routes.SetupRoutes(router)

	// Port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	// Run server
	log.Println("Server run on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed run server: %v", err)
	}
}
