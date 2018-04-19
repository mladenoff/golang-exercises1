package main

import "fmt"

type treeNode struct {
	value string
	left  *treeNode
	right *treeNode
}

func insert(root *treeNode, value string) *treeNode {
	if root == nil {
		return &treeNode{value: value}
	}

	if value < root.value {
		if root.left == nil {
			root.left = &treeNode{value: value}
		} else {
			insert(root.left, value)
		}
	} else {
		if root.right == nil {
			root.right = &treeNode{value: value}
		} else {
			insert(root.right, value)
		}
	}

	return root
}

func traverse(root *treeNode, fn func(*treeNode)) {
	if root == nil {
		return
	}

	if root.left != nil {
		traverse(root.left, fn)
	}

	fn(root)

	if root.right != nil {
		traverse(root.right, fn)
	}
}

func main() {
	var root *treeNode
	root = insert(root, "456")
	root = insert(root, "123")
	root = insert(root, "789")

	traverse(root, func(node *treeNode) {
		fmt.Println(node.value)
	})
}
