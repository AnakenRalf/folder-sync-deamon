package main

import (
	"fmt"
	"folder-sync-deamon/config"
	"folder-sync-deamon/watcher"
	"log"
)

func main() {
	// Load the configuration
	config, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Access the configuration values
	fmt.Printf("Source Path: %s\n", config.SourcePath)
	fmt.Printf("Network Path: %s\n", config.NetworkPath)

	// Initialize and start the watcher
	watcher, err := watcher.NewWatcher(config.SourcePath, config.NetworkPath)
	if err != nil {
		log.Fatalf("Error initializing watcher: %v", err)
	}

	go watcher.Start()

	// Keep the program running
	select {}
}
