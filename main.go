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

	fmt.Printf("Starting the copy process from source: %s to destination: %s\n", folderA, folderB)

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
			fmt.Printf("Creating directory: %s\n", destPath)
			return os.MkdirAll(destPath, info.Mode())
		}

		fmt.Printf("Copying file from: %s to: %s\n", path, destPath)
		return copyFile(path, destPath)
	})

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Copy process completed successfully!")
	}
}

func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("Error opening source file: %s\n", src)
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		fmt.Printf("Error creating destination file: %s\n", dest)
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		fmt.Printf("Error copying from: %s to: %s\n", src, dest)
		return err
	}

	sourceInfo, err := os.Stat(src)
	if err != nil {
		fmt.Printf("Error fetching information for file: %s\n", src)
		return err
	}

	fmt.Printf("Setting permissions for file: %s\n", dest)
	return os.Chmod(dest, sourceInfo.Mode())
}
