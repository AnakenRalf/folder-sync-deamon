package event

import "github.com/fsnotify/fsnotify"

type FileChangeEvent struct {
	Name string
	Op   fsnotify.Op
}

// You can define custom event structures as needed.
