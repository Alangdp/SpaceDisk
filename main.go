package main

import (
	"fmt"
	core "spacedisk/core/Files"
	"spacedisk/types"
	"time"
)

func main() {
	data := &types.DirectoryInfo{
		Filename: "C:",
		Path:     "C:",
		Size:     0,
	}
	root := types.MakeTree(data, nil)

	start := time.Now()
	core.ReadFiles(root, "C://")
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
