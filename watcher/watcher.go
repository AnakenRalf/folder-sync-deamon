package watcher

import (
	"fmt"
	"folder-sync-deamon/sync/network"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	sourcePath  string
	networkPath string
	watcher     *fsnotify.Watcher
}

func NewWatcher(sourcePath, networkPath string) (*Watcher, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	return &Watcher{
		sourcePath:  sourcePath,
		networkPath: networkPath,
		watcher:     watcher,
	}, nil
}

func executeExternalExecutable() error {

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current directory: %v\n", err)
		return err
	}

	fmt.Printf("Current directory: %s\n", currentDir)
	cmd := exec.Command(currentDir + "\\synchronizer.exe")

	// You can set environment variables or provide arguments to the executable if needed
	// cmd.Env = []string{"VAR_NAME=value"}
	// cmd.Args = []string{"arg1", "arg2"}

	cmd.Dir = "." // Set the working directory if necessary

	// Capture the standard output and standard error streams if needed
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Error executing external executable: %v", err)
		return err
	}

	return nil
}

func (w *Watcher) Start() {
	defer w.watcher.Close()

	err := filepath.Walk(w.sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return w.watcher.Add(path)
	})

	if err != nil {
		log.Fatalf("Error adding source folder and subfolders to watcher: %v", err)
	}

	for {
		select {
		case event, ok := <-w.watcher.Events:
			if !ok {
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				// Handle file change events
				log.Printf("File modified: %s", event.Name)

				// Implement synchronization logic with network
				// You can call the synchronization function in the network package here.
				network.SynchronizeWithNetwork(w.sourcePath, w.networkPath, event.Name)
				executeExternalExecutable()
			}

		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Error: %v", err)
		}
	}
}
