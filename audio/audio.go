package audio

import (
	"io"
	"reflect"
	"tracks/checksum"

	"github.com/dhowden/tag"
)

func ParceMetada(file io.ReadSeeker) (string, error){
    // Read the audio file metadata
    metadata, err := tag.ReadFrom(file)
    if err != nil {
        return "",err
    }

	fileCheckSum, err := checksum.SHA1(file) 
	if err != nil {
		return "" ,err
	}

	audioFile := Metadata{
		Format:			string(metadata.Format()),
		FileType:		string(metadata.FileType()),

		Title:			metadata.Title(),
		Album:			metadata.Album(),
		Artist:			metadata.Artist(),
		AlbumArtist:	metadata.AlbumArtist(),
		Composer:		metadata.Composer(),
		Genre:			metadata.Genre(),
		Year:			metadata.Year(),

		// Picture:		metadata.Picture(),
		Lyrics:			metadata.Lyrics(),
		Comment:		metadata.Comment(),

		Raw:			metadata.Raw(),
		Sum:			fileCheckSum,
	}

	return audioFile.ToString(),nil
}


type Metadata struct {
    Format string
	FileType string

	Title string
	Album string
	Artist string
	AlbumArtist string
	Composer string
	Genre string
	Year int

	//Picture *tag.Picture // Artwork
	Lyrics string
	Comment string

	Raw map[string]interface{} // NB: raw tag names are not consistent across formats.

	Sum string
}

// Method with a value receiver
func (metadata Metadata) ToString() (string) {
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