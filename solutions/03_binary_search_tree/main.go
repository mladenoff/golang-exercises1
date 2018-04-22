package main

import (
	"fmt"
	"math/rand"
)

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func insert(root *treeNode, value int) *treeNode {
	if root == nil {
		return &treeNode{value: value}
	}

	if value < root.value {
		root.left = insert(root.left, value)
	} else if root.value < value {
		root.right = insert(root.right, value)
	}

	return root
}

func traverse(root *treeNode, fn func(*treeNode)) {
	if root == nil {
		return
	}

	traverse(root.left, fn)
	fn(root)
	traverse(root.right, fn)
}

func randomInts(targetLength int) []int {
	nums := []int{}
	for len(nums) < targetLength {
		num := rand.Intn(1000)
		nums = append(nums, num)
	}

	return nums
}

func main() {
	var root *treeNode

	for _, value := range randomInts(100) {
		root = insert(root, value)
	}

	traverse(root, func(node *treeNode) {
		fmt.Println(node.value)
	})
}
