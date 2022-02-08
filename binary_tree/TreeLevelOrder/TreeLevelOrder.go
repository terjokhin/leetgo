package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//[3,9,20,null,null,15,7]
	l3_1 := TreeNode{Val: 15, Left: nil, Right: nil}
	l3_2 := TreeNode{Val: 7, Left: nil, Right: nil}
	l2_1 := TreeNode{Val: 9, Left: nil, Right: nil}
	l2_2 := TreeNode{Val: 20, Left: &l3_1, Right: &l3_2}
	l1 := TreeNode{Val: 3, Left: &l2_1, Right: &l2_2}

	fmt.Println(levelOrder(&l1))
}

func levelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	var queue []*TreeNode
	var result [][]int
	var currentLevelData []int
	currentLevelLeftNodesCount := 1
	nextLevelNodesCount := 0

	queue = append(queue, root)

	for len(queue) != 0 {
		if currentLevelLeftNodesCount > 0 {
			node := queue[0]

			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNodesCount++
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNodesCount++
			}

			currentLevelData = append(currentLevelData, node.Val)
			currentLevelLeftNodesCount--
			queue = queue[1:]
		}

		if currentLevelLeftNodesCount == 0 {
			result = append(result, currentLevelData)
			currentLevelLeftNodesCount = nextLevelNodesCount
			nextLevelNodesCount = 0
			currentLevelData = []int{}
		}

	}
	return result
}
