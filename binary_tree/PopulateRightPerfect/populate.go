package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func main() {

	left := Node{Val: 2, Left: nil, Right: nil, Next: nil}
	right := Node{Val: 3, Left: nil, Right: nil, Next: nil}
	root := Node{Val: 1, Left: &left, Right: &right, Next: nil}
	fmt.Println(connect(&root))

	fmt.Println(left)
	fmt.Println(right)

}

func connect(root *Node) *Node {

	if root == nil {
		return nil
	}

	var queue []*Node
	var prevInARow *Node
	currentLevelLeftNodesCount := 1
	nextLevelNodesCount := 0

	queue = append(queue, root)

	for len(queue) != 0 {

		if currentLevelLeftNodesCount > 0 {
			currentNode := queue[0]
			if currentNode.Left != nil {
				nextLevelNodesCount += 2
				queue = append(queue, currentNode.Left, currentNode.Right)
			}

			if prevInARow != nil && currentNode != nil {
				prevInARow.Next = currentNode
			}
			queue = queue[1:]
			prevInARow = currentNode
			currentLevelLeftNodesCount--
		}

		if currentLevelLeftNodesCount == 0 {
			prevInARow = nil
			currentLevelLeftNodesCount = nextLevelNodesCount
			nextLevelNodesCount = 0
		}
	}

	return root
}
