package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/lordofthemind/gormGo/handlers"
	"github.com/lordofthemind/gormGo/initializers"
	"github.com/lordofthemind/gormGo/repositories"
)

func main() {
	// Initialize environment variables, PostgreSQL connection, etc.
	if err := initializers.Initialize(); err != nil {
		log.Fatalf("Failed to initialize: %v", err)
	}

	// Create repository instance
	personRepo := repositories.NewPersonRepository()

	// Create handler instance with injected repository
	personHandler := handlers.NewPersonHandler(personRepo)

	// Initialize Gin router
	router := gin.Default()

	// Register routes
	router.POST("/person", personHandler.CreatePersonHandler)
	router.GET("/person/:id", personHandler.GetPersonByIdHandler)
	router.GET("/persons", personHandler.GetAllPersonsHandler)
	router.PUT("/person/:id", personHandler.UpdatePersonHandler)
	router.DELETE("/person/:id", personHandler.UpdatePersonHandler)

	// Define the port
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090" // Default to port 9090 if PORT environment variable is not set
	}

	// Start the server in a separate goroutine to allow graceful shutdown
	go func() {
		address := fmt.Sprintf(":%s", port)
		if err := router.Run(address); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Listen for termination signals to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Perform any cleanup or finalization tasks here

	log.Println("Server gracefully stopped")
}
