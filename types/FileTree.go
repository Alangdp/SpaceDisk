package types

import (
	"fmt"
	"path/filepath"
	"sort"
	"spacedisk/core"
	"strings"
)

type DirectoryInfo struct {
	Filename string
	Path     string
	Size     int64
	IsFolder bool
}

type Node struct {
	Key      uint32
	Childs   map[uint32]*Node
	Previous *Node
	Data     *DirectoryInfo
}

// PrintDirectoryTree prints the directory tree structure starting from the root node,
// displaying folders and files with their sizes in bytes.
//
// The function prints the tree recursively, using symbols to represent folders (📁)
// and files (📄), and marks the branches with ├── and └── depending on the node position.
//
// Example:
//
//	PrintDirectoryTree(root)
func PrintDirectoryTree(root *Node) {
	var printNode func(node *Node, prefix string, isLast bool)

	printNode = func(node *Node, prefix string, isLast bool) {
		// Determine the symbol for folders and files
		var symbol string
		if node.Data.IsFolder {
			symbol = "📁"
		} else {
			symbol = "📄"
		}

		// Determine branch marker
		var branchMarker string
		if isLast {
			branchMarker = "└── "
		} else {
			branchMarker = "├── "
		}

		// Print current node
		fmt.Printf("%s%s%s %s (Size: %d bytes)\n",
			prefix,
			branchMarker,
			symbol,
			node.Data.Filename,
			node.Data.Size,
		)

		// Prepare children
		if len(node.Childs) > 0 {
			// Sort keys to maintain consistent order
			sortedKeys := make([]uint32, 0, len(node.Childs))
			for k := range node.Childs {
				sortedKeys = append(sortedKeys, k)
			}
			sort.Slice(sortedKeys, func(i, j int) bool {
				return node.Childs[sortedKeys[i]].Data.Filename <
					node.Childs[sortedKeys[j]].Data.Filename
			})

			// Determine new prefix
			var newPrefix string
			if isLast {
				newPrefix = prefix + "    "
			} else {
				newPrefix = prefix + "│   "
			}

			// Print children
			for i, key := range sortedKeys {
				child := node.Childs[key]
				lastChild := (i == len(sortedKeys)-1)
				printNode(child, newPrefix, lastChild)
			}
		}
	}

	// Start printing from the root
	printNode(root, "", true)
}

// getSortedKeys returns the keys of the given map sorted in ascending order.
//
// The function extracts the keys from the map, sorts them, and returns the sorted list.
//
// Example:
//
//	sortedKeys := getSortedKeys(nodes)
//	fmt.Println(sortedKeys)
func GetSortedKeys(nodes map[uint32]*Node) []uint32 {
	keys := make([]uint32, 0, len(nodes))
	for k := range nodes {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return keys
}

// getFileTypeExtension returns the file extension of the given filename,
// formatted with gray color escape codes for terminal output.
//
// If the filename has an extension, it returns it in gray; otherwise, returns an empty string.
//
// Example:
//
//	ext := getFileTypeExtension("file.txt")
//	fmt.Println(ext) // Output: "\033[90m.txt\033[0m"
func GetFileTypeExtension(filename string) string {
	ext := filepath.Ext(filename)
	if ext != "" {
		return "\033[90m" + ext + "\033[0m" // Gray color for extension
	}
	return ""
}

// MakeEmptyTree, return a empty Tree
// Example:
//
//	MakeEmptyTree()
func MakeEmptyTree() *Node {
	return &Node{}
}

// MakeTree creates a tree structure with the given data,
// optionally allowing a previous node to be set as its parent.
//
// The function returns a new tree node. If a previous node is provided,
// it will be linked as the parent of the new node.
//
// Example:
//
//	MakeTree(data, previousNode) // Where 'previousNode' is optional
func MakeTree(data *DirectoryInfo, previous *Node) *Node {
	hashedKey := core.Fnv1aHash(data.Filename)

	return &Node{
		Key:      hashedKey,
		Data:     data,
		Childs:   make(map[uint32]*Node),
		Previous: previous,
	}
}

// AppendChild adds a new child node to the root node, using the provided data
// to create the new node. If the root node's children map is not initialized,
// it initializes it before adding the new node.
//
// The function returns the newly created child node.
//
// Example:
//
//	newNode := AppendChild(root, data)
//	fmt.Println("New node added:", newNode)
func AppendChild(root *Node, data *DirectoryInfo) *Node {
	// Make a new node with provided data
	newNode := MakeTree(data, root)

	// Validate if root have childs map
	if root.Childs == nil {
		root.Childs = make(map[uint32]*Node)
	}

	root.Childs[newNode.Key] = newNode

	return newNode
}

// AppendFullPath adds nodes representing directories and files to the tree,
// based on the provided path. It traverses the path, and for each segment,
// it checks if the node already exists. If not, it creates a new node and
// adds it to the parent node. The function returns the root node of the tree.
//
// The function processes the path segment by segment, creating a new node
// for each directory or file that does not already exist in the tree.
//
// Example:
//
//	AppendFullPath(root, "path/to/directory")
//	fmt.Println("Updated tree with new nodes")
func AppendFullPath(root *Node, data *DirectoryInfo) *Node {
	pathFormatted := filepath.ToSlash(data.Path)
	pathSplited := strings.Split(pathFormatted, "/")

	temp := root
	currentPath := ""

	for i, segment := range pathSplited {
		// Construir o caminho acumulativo
		if i > 0 {
			currentPath += "/"
		}
		currentPath += segment

		// Calcular hash
		hash := core.Fnv1aHash(strings.ToLower(segment))

		// Criar nó, se não existir
		if _, exists := temp.Childs[hash]; !exists {
			newNode := &Node{
				Data: &DirectoryInfo{
					Filename: segment,
					Path:     currentPath,
					Size:     data.Size,
					IsFolder: data.IsFolder,
				},
				Childs: make(map[uint32]*Node),
			}
			temp.Childs[hash] = newNode
		}

		// Avançar para o próximo nó
		temp = temp.Childs[hash]
	}

	return root
}
