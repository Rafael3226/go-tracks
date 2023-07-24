package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"tracks/files"
)

func main() {
    // Use filepath.Join to create the path that works on both Linux and Windows
    // For Windows: "D:\Music", For Linux: "D:/Music"
    folderPath := filepath.Join("D:", "Music")

    metadataList, failedPaths, err := files.ListFolderContent(folderPath)
    if err != nil {
        // log.Fatalf("Error Listing Folder Content: %v", err)
    }

    metadataContent := strings.Join(metadataList,"\n")
    metadataCSV := "all_tracks.csv"
    CreateFile(metadataCSV,metadataContent)

    faildContent := strings.Join(failedPaths,"\n")
    failedCSV := "failed.csv"
    CreateFile(failedCSV,faildContent)
}

func CreateFile(name string, content string) {
    file, err := os.Create(name)

    _, err = file.WriteString(content)
    if err != nil {
        log.Fatalf("Error writing to file: %v", err)
    }
}
