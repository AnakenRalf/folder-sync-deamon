package sync

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Synchronize(sourcePath, networkPath string) {
	// List all files in the source folder
	err := filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error walking source directory: %v", err)
			return err
		}

		if !info.IsDir() {
			relativePath, _ := filepath.Rel(sourcePath, path)
			destinationPath := filepath.Join(networkPath, relativePath)

			// Copy or synchronize the file to the network folder
			err := copyFile(path, destinationPath)
			if err != nil {
				log.Printf("Error copying %s to network folder: %v", path, err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error synchronizing files: %v", err)
	}
	fmt.Println("Synchronization complete.")
}

func copyFile(sourcePath, destinationPath string) error {
	// Ensure that the destination directory path exists
	destinationDir := filepath.Dir(destinationPath)
	if err := os.MkdirAll(destinationDir, os.ModePerm); err != nil {
		return err
	}

	// Open the source file for reading
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create the destination file for writing
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Copy the content from the source file to the destination file
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	// Ensure that all data is flushed to the destination file
	err = destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}
