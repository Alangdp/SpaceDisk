package core

import (
	"fmt"
	"os"
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
		filePath := rootPath + "/" + file.Name()

		if file.IsDir() {
			stringArray = append(stringArray, ReadFiles(filePath)...)
		} else {
			_, err := os.Stat(filePath)
			if err != nil {
				fmt.Println("Error reading file info", rootPath)
				return stringArray
			}
		}

		stringArray = append(stringArray, []string{rootPath, file.Name()})
	}

	return stringArray
}