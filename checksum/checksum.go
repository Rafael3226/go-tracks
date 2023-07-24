package checksum

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func SHA1(file io.Reader) (string,error){
	// Calculate the SHA1 checksum of the audio data
	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "",err
	}

	// Get the SHA1 checksum as a byte slice
	checksumBytes := hash.Sum(nil)

	// Convert the byte slice to a hexadecimal string
	checksum := fmt.Sprintf("%x", checksumBytes)

	return checksum,nil
}