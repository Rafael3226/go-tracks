package files

import (
	"os"
	"path/filepath"
)


func ListFolderContent(folderPath string) ([]string, error) {
    var fileList []string

    // Walk the folder and list its contents
    err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        // Skip the folder itself
        if path == folderPath {
            return nil
        }

        fileList = append(fileList, path)
        return nil
    })

    if err != nil {
        return nil, err
    }

    return fileList, nil
}
