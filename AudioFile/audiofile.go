package audiofile

import (
	"tracks/checksum"
	"github.com/dhowden/tag"
)

func ParceMetada(file) (Metadata, error){
    // Read the audio file metadata
    metadata, err := tag.ReadFrom(file)
    if err != nil {
        return nil,error
    }

	audioFile := Metadata{
		Format:			metadata.Format(),
		FileType:		metadata.FileType(),

		Title:			metadata.Title(),
		Album:			metadata.Album(),
		Artist:			metadata.Artist(),
		AlbumArtist:	metadata.AlbumArtist(),
		Composer:		metadata.Composer(),
		Genre:			metadata.Genre(),
		Year:			metadata.Year(),

		Picture:		metadata.Picture(),
		Lyrics:			metadata.Lyrics(),
		Comment:		metadata.Comment(),

		Raw:			metadata.Raw(),
		Sum:			checksum.SHA1(file),
	}

	return audioFile,nil
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

	Picture *Picture // Artwork
	Lyrics string
	Comment string

	Raw map[string]interface{} // NB: raw tag names are not consistent across formats.

	Sum string
}