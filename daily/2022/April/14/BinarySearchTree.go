package main

import "log"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	t5 := TreeNode{3, nil, nil}
	t4 := TreeNode{1, nil, nil}
	t3 := TreeNode{2, &t4, &t5}
	t2 := TreeNode{7, nil, nil}
	t1 := TreeNode{4, &t3, &t2}

	log.Println(searchBST(&t1, 2))

}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	current := root.Val

	if current == val {
		return root
	}

	var result *TreeNode

	if current > val {
		result = searchBST(root.Left, val)
	}

	if current < val {
		result = searchBST(root.Right, val)
	}

	return result

}
