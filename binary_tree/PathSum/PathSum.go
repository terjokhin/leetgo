package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	l2 := TreeNode{Val: 2, Left: nil, Right: nil}
	l1 := TreeNode{Val: 1, Left: &l2, Right: nil}

	fmt.Println(hasPathSum(&l1, 1))
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root != nil {
		return pathSum(root, 0, targetSum)
	}
	return false
}

func pathSum(node *TreeNode, currentSum, target int) bool {
	if node.Left == nil && node.Right == nil && currentSum+node.Val == target {
		return true
	} else {

		left := false
		if node.Left != nil {
			left = pathSum(node.Left, currentSum+node.Val, target)
		}

		right := false
		if node.Right != nil {
			right = pathSum(node.Right, currentSum+node.Val, target)
		}

		return right || left
	}
}
