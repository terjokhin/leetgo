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

	fmt.Println(buildTree([]int{1, 2}, []int{2, 1}))
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	pLen := len(preorder)

	if pLen == 0 {
		return nil
	}

	head := preorder[0]

	if pLen == 1 {
		return &TreeNode{Val: head, Left: nil, Right: nil}
	}

	idx := -1

	for i, v := range inorder {
		if v == head {
			idx = i
			break
		}
	}

	tail := preorder[1:]

	left := buildTree(tail[:idx], inorder[0:idx])
	right := buildTree(tail[idx:], inorder[idx+1:])

	return &TreeNode{Val: head, Left: left, Right: right}
}
