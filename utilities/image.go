package utilities

import (
	"encoding/base64"
	"os"
)

func Base64ToImage(base64String, filename string) error {
	// decode a base64 string
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return err
	}
	// create a image file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// write data in file
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
