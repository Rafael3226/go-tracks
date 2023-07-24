package main

import (
	"log"
	"os"
	"path/filepath"
	"tracks/audiofile"
	"tracks/files"
)

func main() {
    // Use filepath.Join to create the path that works on both Linux and Windows
    // For Windows: "D:\Music", For Linux: "D:/Music"
    folderPath := filepath.Join("D:", "Music") 
    

    pathList, err := files.ListFolderContent(folderPath)
    if err != nil {
        log.Fatalf("Error Listing Folder Content: %v", err)
    }

    for _ , path := range pathList {
        // Open the audio file
        file, err := os.Open(path)
        if err != nil {
            log.Fatalf("Error opening the audio file: %v", err)
        }
        defer file.Close()

        metadata,err := audiofile.ParceMetada(file)
        if err != nil {
            log.Fatalf("Error Parcing the Metadata: %v", err)
        }
        
    }

    content := ""
    
    newCSV := "all_tracks.csv"
    file, err := os.Create(newCSV)

    _, err = file.WriteString(content)
    if err != nil {
        log.Fatalf("Error writing to file: %v", err)
    }
}