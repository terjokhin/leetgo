package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//[1,null,2,3]
	l3 := TreeNode{Val: 3, Left: nil, Right: nil}
	l2 := TreeNode{Val: 2, Left: &l3, Right: nil}
	l1 := TreeNode{Val: 1, Left: nil, Right: &l2}

	fmt.Println(inorderTraversal(&l1))
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	current := []int{root.Val}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	return append(left, append(current, right...)...)
}
