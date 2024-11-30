package types

import (
	"fmt"
	"spacedisk/core"
)

type DirectoryInfo struct {
	filename string
	path     string
	size     int
}

type Node struct {
	key      uint32
	childs   map[int32]*Node
	previous *Node
}

func main() {

	root := &Node{
		key: uint32(core.Fnv1aHash("C:/")),
		childs: make(map[int32]*Node),
	}

	fmt.Println(root)
}
