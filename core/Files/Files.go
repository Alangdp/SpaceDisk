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

	// Percorrer os arquivos e diretórios encontrados
	for _, file := range files {
		filePath := filepath.Join(rootPath, file.Name())

		if file.IsDir() {
			// Adiciona o diretório à árvore e continua a exploração recursiva
			newNode := types.AppendFullPath(root, filePath)
			// Recursivamente lê os arquivos do diretório
			ReadFiles(newNode, filePath)
		} else {
			// Adiciona o arquivo à árvore
			types.AppendFullPath(root, filePath)
		}
	}

	return root
}
