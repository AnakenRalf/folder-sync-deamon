package main

import (
	"folder-sync-deamon/config"
	"folder-sync-deamon/sync"
	"log"
)

func main() {
	// Load the configuration
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Access the configuration values
	sourcePath := config.SourcePath
	networkPath := config.NetworkPath

	// Initialize and run the synchronization logic
	sync.Synchronize(sourcePath, networkPath)
}
