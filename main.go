package main

// Import type from Types/BinaryTree.go

import (
	"fmt"
	"os"
	"time"
)

// ReadFiles return all files from rootPath recursively.
//
// Example case use:
//
//	ReadFiles("C:/")
func ReadFiles(rootPath string) [][]string {
	stringArray := [][]string{}

	files, err := os.ReadDir(rootPath)

	if err != nil {
		fmt.Println("Error reading file info", rootPath)
		return stringArray
	}

	for _, file := range files {
		if file.IsDir() {
			stringArray = append(stringArray, ReadFiles(rootPath+"/"+file.Name())...)
		} else {
			_, err := os.Stat(rootPath + "/" + file.Name())
			if err != nil {
				fmt.Println("Error reading file info", rootPath)
				return stringArray
			}
		}

		stringArray = append(stringArray, []string{rootPath, file.Name()})
	}

	return stringArray
}

func main() {
	startTime := time.Now()
	files := ReadFiles("C:/")
	fmt.Println("Time taken: ", time.Since(startTime))
	fmt.Println("Files: ", len(files))
}
