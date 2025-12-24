package main

import (
	"log"
)

func main() {
	log.Println("Starting frontrun-monitor backend...")

	// Load configuration
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize server
	server := NewServer(cfg)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
