package services

import (
	"errors"
	"os"
)

func GetFilePath(fileName string) (string, error) {
	basePath := "data"

	filePath := basePath + "/" + fileName
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.New("file not found")
	}

	return filePath, nil
}
