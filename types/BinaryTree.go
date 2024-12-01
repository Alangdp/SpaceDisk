package types

import "fmt"

type BynaryTree struct {
	Value int8
	Left  *BynaryTree
	Right *BynaryTree
}

// Inorder_traversal function
// Inorder traversal is a type of depth-first traversal that visits the left branch, then the current node, and finally, the right branch of a tree.
func Inorder_traversal(root *BynaryTree) {
	if root != nil {
		Inorder_traversal(root.Left)
		fmt.Println(root.Value)
		Inorder_traversal(root.Right)
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
