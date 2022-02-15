package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	right := TreeNode{Val: 2, Left: nil, Right: nil}
	root := TreeNode{Val: 1, Left: nil, Right: &right}
	fmt.Println(lowestCommonAncestor(&root, &root, &right))

}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var lca TreeNode
	goDeeper(root, p, q, &lca)
	return &lca
}

func goDeeper(current, p, q, result *TreeNode) bool {
	if current == nil {
		return false
	}

	this := current == p || current == q
	left := goDeeper(current.Left, p, q, result)
	right := goDeeper(current.Right, p, q, result)

	if this && left || this && right || left && right {
		*result = *current
		return true
	}

	return this || left || right
}
