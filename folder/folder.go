package folder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"tracks/file"
	"tracks/object"
)

func ListContent(folderPath string) ([]string, error) {
	var fileList []string
	fileList = append(fileList, "Path,Title,Extension,Sum")
	// Walk the folder and list its contents
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip the folder itself
		if path == folderPath {
			return nil
		}

		file, err := file.NewFile(path)
		if err != nil {
			return err
		}
		if file != nil {
			obj := object.New(file)
			fileData := strings.Join(obj.ParceValuesToStringList(), ",")
			fileList = append(fileList, fileData)
		}

		fmt.Printf("\rFile_#_%v", len(fileList))
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}
