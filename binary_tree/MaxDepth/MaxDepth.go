package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	l3_1 := TreeNode{Val: 15, Left: nil, Right: nil}
	l3_2 := TreeNode{Val: 7, Left: nil, Right: nil}
	l2_1 := TreeNode{Val: 9, Left: nil, Right: nil}
	l2_2 := TreeNode{Val: 20, Left: &l3_1, Right: &l3_2}
	l1 := TreeNode{Val: 3, Left: &l2_1, Right: &l2_2}

	fmt.Println(maxDepth(&l1))
}

func maxDepth(root *TreeNode) int {
	if root != nil {
		return traverseDeeper(root, 0)
	}

	return 0
}

func traverseDeeper(node *TreeNode, currentMax int) int {
	if node == nil {
		return currentMax
	}

	left := traverseDeeper(node.Left, currentMax+1)
	right := traverseDeeper(node.Right, currentMax+1)

	if left > right {
		return left
	} else {
		return right
	}
}
