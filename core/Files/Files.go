package core

import (
	"log"
	"os"
	"path/filepath"
	"spacedisk/types"
)

// ReadFiles returns all files from the given rootPath recursively, constructing a tree of directories and files.
//
// This function reads the contents of the specified directory and its subdirectories,
// adding each directory and file to a tree structure, starting from the provided root node.
//
// Example usage:
//
//	ReadFiles("C:/")
func ReadFiles(root *types.Node, rootPath string) *types.Node {
	files, err := os.ReadDir(rootPath)
	if err != nil {
		log.Printf("Error reading directory %s: %v", rootPath, err)
		return root
	}

	for _, file := range files {
		filePath := filepath.Join(rootPath, file.Name())

		// Adicionar informações do arquivo/diretório ao nó
		info, _ := file.Info()
		data := &types.DirectoryInfo{
			Filename: info.Name(),
			Path:     filePath,
			Size:     func() int64 {
				if file.IsDir() {
					return DirSizeMB(filePath)
				}
				return info.Size()
			}(),
			IsFolder: file.IsDir(),
		}

		types.AppendFullPath(root, data)
		// Recursão para diretórios

		if file.IsDir() {
			ReadFiles(root, filePath)
		} 
	}

	return root
}


// DirSizeMB calculates the total size of all files in a given directory (in bytes).
// It walks through the directory tree and sums up the sizes of all files found.
// Note: This function ignores any errors encountered while traversing the directory.
//
// Reference: https://stackoverflow.com/questions/32482673/how-to-get-directory-total-size
//
// Parameters:
// - path: The path to the directory whose size is to be calculated.
//
// Returns:
// - int64: The total size of all files in the directory, in bytes.
func DirSizeMB(path string) int64 {
    sizes := make(chan int64)
    readSize := func(path string, file os.FileInfo, err error) error {
        if err != nil || file == nil {
            return nil // Ignore errors
        }
        if !file.IsDir() {
            sizes <- file.Size()
        }
        return nil
    }

    go func() {
        filepath.Walk(path, readSize)
        close(sizes)
    }()

    size := int64(0)
    for s := range sizes {
        size += s
    }

    return size
}
