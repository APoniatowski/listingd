package main

import (
	"log"
	"net/http"

	"github.com/APoniatowski/listingd/config"
	"github.com/APoniatowski/listingd/pkg"
	"github.com/APoniatowski/listingd/pkg/kafka"
)

func main() {
	// Load the configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to the database
	db := pkg.SetupDatabase(cfg)

	// Initialize Gin router
	router := pkg.SetupRouter(db)

	kafkaProducer, err := kafka.NewProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer kafkaProducer.Close()

	// Start the server
	if err := router.Run(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting the server: %v", err)
	}
}
