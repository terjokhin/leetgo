package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	l5 := TreeNode{Val: 3, Left: nil, Right: nil}
	l4 := TreeNode{Val: 3, Left: nil, Right: nil}
	l3 := TreeNode{Val: 2, Left: nil, Right: &l5}
	l2 := TreeNode{Val: 2, Left: nil, Right: &l4}
	l1 := TreeNode{Val: 1, Left: &l2, Right: &l3}

	fmt.Println(isSymmetric(&l1))
}

func isSymmetric(root *TreeNode) bool {
	if root != nil {
		return isMirror(root.Left, root.Right)
	}
	return true
}

func isMirror(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left != nil && right != nil && left.Val == right.Val {
		return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}
	return false
}
