package utils

import (
	"os"
	"path/filepath"
)

func GetCurrentBinaryDir(otherCurrentDir string) string {
	var dir string

	if otherCurrentDir == "" {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	} else {
		dir, _ = filepath.Abs(filepath.Dir(otherCurrentDir))
	}

	return dir
}

func GetPathAbsValidated(otherCurrentDir string, fileName string) string {
	currentDir := GetCurrentBinaryDir(otherCurrentDir)
	return filepath.Join(currentDir, fileName)
}

func IsFileExist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}
