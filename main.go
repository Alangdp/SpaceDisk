package main

import (
	core "spacedisk/core/Files"
	"spacedisk/types"
)

func main() {
	data := &types.DirectoryInfo{
		Filename: "C:",
		Path:     "C:",
		Size:     0,
	}
	root := types.MakeTree(data, nil)

	newRoot := core.ReadFiles(root, "E:\\Jogos  xbox\\Planet of Lana")

	types.PrintDirectoryTree(newRoot)

	// data := &types.DirectoryInfo{
	// 	Filename: "C:",
	// 	Path:     "C:",
	// 	Size:     0,
	// }
	// root := types.MakeTree(data, nil)

	// // Nó para "Users"
	// usersData := &types.DirectoryInfo{
	// 	Filename: "Users",
	// 	Path:     "C:/Users/",
	// 	Size:     0,
	// }
	// usersNode := types.MakeTree(usersData, root)
	// root.Childs[usersNode.Key] = usersNode

	// // Nó para "gabri"
	// gabriData := &types.DirectoryInfo{
	// 	Filename: "gabri",
	// 	Path:     "C:/Users/gabri/",
	// 	Size:     0,
	// }
	// gabriNode := types.MakeTree(gabriData, usersNode)
	// usersNode.Childs[gabriNode.Key] = gabriNode

	// // Nó para "Pictures"
	// picturesData := &types.DirectoryInfo{
	// 	Filename: "Pictures",
	// 	Path:     "C:/Users/gabri/Pictures/",
	// 	Size:     0,
	// }
	// picturesNode := types.MakeTree(picturesData, gabriNode)
	// gabriNode.Childs[picturesNode.Key] = picturesNode

	// // Nó para "Feedback"
	// feedbackData := &types.DirectoryInfo{
	// 	Filename: "Feedback",
	// 	Path:     "C:/Users/gabri/Pictures/Feedback/",
	// 	Size:     0,
	// }
	// feedbackNode := types.MakeTree(feedbackData, picturesNode)
	// picturesNode.Childs[feedbackNode.Key] = feedbackNode

	// // Nó para "{F8D3D7C3-23F8-4DF3-B34A-294D63831505}"
	// finalData := &types.DirectoryInfo{
	// 	Filename: "{F8D3D7C3-23F8-4DF3-B34A-294D63831505}",
	// 	Path:     "C:/Users/gabri/Pictures/Feedback/{F8D3D7C3-23F8-4DF3-B34A-294D63831505}",
	// 	Size:     0,
	// }
	// finalNode := types.MakeTree(finalData, feedbackNode)
	// feedbackNode.Childs[finalNode.Key] = finalNode

	// types.PrintDirectoryTree(root)
}
