package chefgi

import (
	"os"
	"path/filepath"
)

func GetFilePath() (string, error) {
	rootPath, err := os.Getwd()

	if err != nil {
		return "", err
	}

	file := os.Args[1]

	return filepath.Join(rootPath, file), nil
}