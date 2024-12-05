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
			Size:     info.Size(),
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
