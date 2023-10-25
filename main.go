package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	folderA := "./folderA" // Source folder
	folderB := "./folderB" // Destination folder

	err := filepath.Walk(folderA, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(folderA, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(folderB, relPath)

		if info.IsDir() {
			return os.MkdirAll(destPath, info.Mode())
		}

		return copyFile(path, destPath)
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dest, sourceInfo.Mode())
}
