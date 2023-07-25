package io

import (
	"io"
	"log"
	"os"
)

func MoveFile(oldPath string, newPath string) error {
	if oldPath == newPath {
		return nil
	}
	fileExist, err := isFileExists(newPath)
	if err != nil {
		return err
	}
	if !fileExist {
		copyFile(oldPath, newPath)
	}
	return nil
}

func copyFile(sourcePath, destinationPath string) error {
	// Open the source file for reading
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	// Create or truncate the destination file
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// Copy the content from the source to the destination
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func isFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil // File exists
	} else if os.IsNotExist(err) {
		return false, nil // File does not exist
	}
	// There was an error accessing the file (e.g., permission issue)
	return false, err
}

func IsDir(path string) bool {
	// Get information about the file/directory at the specified fileFactory
	fileInfo, err := os.Stat(path)
	if err != nil {
		return true
	}

	// Check if the fileFactory is a directory
	return fileInfo.IsDir()
}

func CreateFile(name string, content string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatalf("Error creating the file: %v", err)
	}
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}
