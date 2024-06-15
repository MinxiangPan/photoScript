package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputPath := os.Args[1]
	targetFileExtension := ".hif"
	processFileExtension := ".arw"

	if len(os.Args) > 2 {
		targetFileExtension = os.Args[2]
	}
	fmt.Printf("use %s target extension\n", targetFileExtension)

	if len(os.Args) > 3 {
		processFileExtension = os.Args[3]
	}
	fmt.Printf("use %s process extension\n", processFileExtension)

	fmt.Printf("reading files from path: %s\n", inputPath)
	files, err := os.ReadDir(inputPath)
	if err != nil {
		fmt.Println("unable to read files from dir")
		return
	}

	targetFiles := make(map[string]bool)
	deleteArwFiles := make([]string, 0)
	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(strings.ToLower(fileName), targetFileExtension) {
			targetFiles[fileName[:len(fileName)-len(targetFileExtension)]] = true
		}
	}

	for _, file := range files {
		fileName := file.Name()
		if strings.HasSuffix(strings.ToLower(fileName), processFileExtension) {
			fileNameWOSuffix := fileName[:len(fileName)-len(processFileExtension)]
			if !targetFiles[fileNameWOSuffix] {
				fileFullPath := filepath.Join(inputPath, fileName)
				deleteArwFiles = append(deleteArwFiles, fileFullPath)
				os.Remove(fileFullPath)
			}
		}
	}

	fmt.Printf("delete %d %s file\n here are the list", len(deleteArwFiles), processFileExtension)
	fmt.Println(deleteArwFiles)
}
