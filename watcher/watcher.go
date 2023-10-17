package watcher

import (
	"folder-sync-deamon/sync/network"
	"log"
	"os"
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
			}

		case err, ok := <-w.watcher.Errors:
			if !ok {
				return
			}
			log.Printf("Error: %v", err)
		}
	}
}
