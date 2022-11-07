package utils

import (
	"log"
	"os"
)

func CreateDirIfNotExists(dirName string) {
	_, err := os.Stat(dirName)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}

func TruncateFile(filePath string) error {
	err := os.Truncate(filePath, 200)
	if err != nil {
		return err
	}

	return nil
}
