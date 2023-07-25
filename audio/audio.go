package audio

import (
	"fmt"
	"io"
	"reflect"
	"tracks/checksum"

	"github.com/dhowden/tag"
)

func ParceMetada(file io.ReadSeeker) (string, error) {
	// Read the audio file metadata
	tags, err := tag.ReadFrom(file)
	if err != nil {
		return "", err
	}

	fileCheckSum, err := checksum.SHA1(file)
	if err != nil {
		return "", err
	}

	metadata := Metadata{
		Format:   string(tags.Format()),
		FileType: string(tags.FileType()),

		Title:       tags.Title(),
		Album:       tags.Album(),
		Artist:      tags.Artist(),
		AlbumArtist: tags.AlbumArtist(),
		Composer:    tags.Composer(),
		Genre:       tags.Genre(),
		Year:        fmt.Sprint(tags.Year()),

		// Picture:		tags.Picture(),
		Lyrics:  tags.Lyrics(),
		Comment: tags.Comment(),

		// Raw:			tags.Raw(),
		Sum: fileCheckSum,
	}

	return metadata.ToString(), nil
}

type Metadata struct {
	Format   string
	FileType string

	Title       string
	Album       string
	Artist      string
	AlbumArtist string
	Composer    string
	Genre       string
	Year        string

	//Picture *tag.Picture // Artwork
	Lyrics  string
	Comment string

	// Raw map[string]interface{} // NB: raw tag names are not consistent across formats.

	Sum string
}

// Method with a value receiver
func (metadata Metadata) ToString() string {
	var str string

	// Get the reflect.Value of the struct
	v := reflect.ValueOf(metadata)

	// Check if the value is a struct
	if v.Kind() == reflect.Struct {
		// Loop through the fields of the struct
		for i := 0; i < v.NumField(); i++ {
			// Get the field name and value
			//fieldName := v.Type().Field(i).Name
			fieldValue := v.Field(i)
			str += fieldValue.String() + ","
		}
	}
	return str
}
