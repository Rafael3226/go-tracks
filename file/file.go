package file

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"tracks/checksum"
	"tracks/io"
)

type FileFactory struct {
	Path string
}

type File struct {
	FileFactory string
	Title       string
	Extension   string
	Sum         string
}

func NewFile(dir string) (*File, error) {
	if io.IsDir(dir) {
		return nil, nil
	}

	fileFactory := newFileFactory(dir)
	err := io.MoveFile(dir, fileFactory.Path)
	if err != nil {
		return nil, err
	}

	return fileFactory.createFile()
}

func newFileFactory(dir string) *FileFactory {
	path := fixPath(dir)
	return &FileFactory{
		Path: path,
	}
}

func fixPath(path string) string {
	newPath := filepath.Base(path)
	newPath = strings.Replace(newPath, ",", "&", -1)

	removeCamelot := regexp.MustCompile("(?:[1-9]|1[0-2])[AB] - (?:1[0-9][0-9]|[1-9][0-9])")
	newPath = removeCamelot.ReplaceAllString(newPath, "")

	removeHyphen := regexp.MustCompile(`^[0-9\s.\-_]*`)
	newPath = removeHyphen.ReplaceAllString(newPath, "")

	removeDoubleSpaces := regexp.MustCompile(` {2,}`)
	newPath = removeDoubleSpaces.ReplaceAllString(newPath, " ")

	newPath = strings.TrimRight(newPath, " ")

	fullPath := filepath.Join("D:", "TRACKS", newPath)
	return fullPath
}

func (fileFactory *FileFactory) getTitle() string {
	return filepath.Base(fileFactory.Path)
}

func (fileFactory *FileFactory) getExtension() string {
	return filepath.Ext(fileFactory.Path)
}

func (fileFactory *FileFactory) getCheckSum() (string, error) {
	file, err := os.Open(fileFactory.Path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	return checksum.SHA1(file)
}

func (fileFactory *FileFactory) createFile() (*File, error) {
	sum, err := fileFactory.getCheckSum()
	if err != nil {
		return nil, err
	}

	return &File{
		FileFactory: fileFactory.Path,
		Title:       fileFactory.getTitle(),
		Extension:   fileFactory.getExtension(),
		Sum:         sum,
	}, nil
}
