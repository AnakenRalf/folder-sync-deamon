package network

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func SynchronizeWithNetwork(sourcePath, networkPath, filePath string) error {
	// Build the source file and destination file paths
	sourceFile := filepath.Join(sourcePath, filePath)
	destinationFile := filepath.Join(networkPath, filePath)

	// Open the source file for reading
	source, err := os.Open(sourceFile)
	if err != nil {
		return err
	}
	defer source.Close()

	// Create the destination file
	destination, err := os.Create(destinationFile)
	if err != nil {
		return err
	}
	defer destination.Close()

	// Copy the content from the source file to the destination file
	_, err = io.Copy(destination, source)
	if err != nil {
		return err
	}

	fmt.Printf("Synchronized: %s\n", filePath)
	return nil
}
