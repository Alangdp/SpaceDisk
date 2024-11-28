package main

// Import type from Types/BinaryTree.go

import (
	"fmt"
	"os"
	"time"
)

func readFiles(rootPath string) [][]string {
	stringArray := [][]string{}

	files, err := os.ReadDir(rootPath)

	if err != nil {
		fmt.Println("Error reading file info", rootPath)
		return stringArray
	}

	for _, file := range files {
		if file.IsDir() {
			// fmt.Println("Directory: ", rootPath)
			stringArray = append(stringArray, readFiles(rootPath+"/"+file.Name())...)
		} else {
			_, err := os.Stat(rootPath + "/" + file.Name())
			if err != nil {
				fmt.Println("Error reading file info", rootPath)
				return stringArray
			}

			// fmt.Println("File: ", fileInfo.Name(), fileInfo.Size())
		}

		stringArray = append(stringArray, []string{rootPath, file.Name()})
	}

	return stringArray
}

func main() {
	startTime := time.Now()
	files := readFiles("C:/")
	fmt.Println("Time taken: ", time.Since(startTime))
	fmt.Println("Files: ", len(files))
}
