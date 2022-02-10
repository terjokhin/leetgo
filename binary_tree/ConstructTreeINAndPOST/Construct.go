package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	fmt.Println(buildTree([]int{2, 1}, []int{2, 1}))

}

func buildTree(inorder []int, postorder []int) *TreeNode {
	pLen := len(postorder)

	if pLen == 0 {
		return nil
	}

	if pLen == 1 {
		return &TreeNode{Val: postorder[0], Left: nil, Right: nil}
	}

	value := postorder[pLen-1]
	idx := pLen

	for i, v := range inorder {
		if v == value {
			idx = i
			break
		}
	}

	left := buildTree(inorder[0:idx], postorder[0:idx])
	right := buildTree(inorder[idx+1:], postorder[idx:pLen-1])

	return &TreeNode{
		Val:   value,
		Left:  left,
		Right: right,
	}
}
