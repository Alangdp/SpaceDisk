package main

import "fmt"

type BynaryTree struct {
	Value int8
	Left  *BynaryTree
	Right *BynaryTree
}

func inorder_traversal(root *BynaryTree) {
	if root != nil {
		inorder_traversal(root.Left)
		fmt.Println(root.Value)
		inorder_traversal(root.Right)
	}
}

func Insert(root *BynaryTree, data int8) *BynaryTree {
	temp := &BynaryTree{
		Value: data,
		Left:  nil,
		Right: nil,
	}

	if root == nil {
		root = temp
		return temp
	} else {
		current := root
		var parent *BynaryTree

		for {
			parent = current
			if data < parent.Value {
				current = current.Left

				if current == nil {
					parent.Left = temp
					return root
				}
			} else {
				current = current.Right

				if current == nil {
					parent.Right = temp
					return root
				}
			}
		}
	}
}

func main() {
	array := [15]int8{34, 84, 15, 0, 2, 99, 79, 9, 88, 89, 18, 31, 39, 100, 101}
	var root *BynaryTree

	for i := 0; i < len(array); i++ {
		root = Insert(root, array[i])
	}

	inorder_traversal(root)
}
