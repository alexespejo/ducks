package ioFile

import (
	"fmt"
	"os"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

func CreateFile(filePath string) error {
	if !FileExists(filePath) {
		file, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer file.Close()
	}

	return nil
}

func WriteToFile(route string, content string) error {
	if !FileExists(route) {
		return fmt.Errorf("file does not exist at path: %s", route)
	}

	return os.WriteFile(route, []byte(content), 0644)
}
