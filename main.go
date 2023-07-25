package main

import (
	"log"
	"path/filepath"
	"strings"
	"tracks/folder"
	"tracks/io"
)

func main() {
	GetDataFromFolder()
}

func GetDataFromFolder() {
	// Path
	folderPath := filepath.Join("D:", "Music")
	//
	metadataList, err := folder.ListContent(folderPath)
	if err != nil {
		log.Fatalf("Error Listing Folder Content: %v", err)
	}
	// Tags CSV
	metadataContent := strings.Join(metadataList, "\n")
	metadataCSV := "all_tracks.csv"
	io.CreateFile(metadataCSV, metadataContent)
}
