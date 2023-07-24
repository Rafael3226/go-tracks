package files

import (
	"io"
	"os"
	"path/filepath"
	"tracks/audio"
)


func ListFolderContent(folderPath string) ([]string, []string, error) {
    var fileList []string
    var failedPaths []string

    // Walk the folder and list its contents
    err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        // Skip the folder itself
        if path == folderPath {
            return nil
        }

        fileData, err := ToString(path)
        if err != nil {
            failedPaths = append(failedPaths, path)
        }
        fileList = append(fileList, fileData)
        return nil
    })

    if err != nil {
        return nil,nil, err
    }

    return fileList, failedPaths, nil
}

func ToString(path string) (string, error) {
    // Open the audio file
    file, err := os.Open(path)
    if err != nil {
        return "",err
    }
    defer file.Close()

    // Convert *os.File to io.ReadSeeker
	var readSeeker io.ReadSeeker = file
    metadata,err := audio.ParceMetada(readSeeker)
    if err != nil {
        return "",err
    }

    return metadata,nil
}