package main

import "fmt"

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {

	n1 := Node{2, nil}
	n := Node{1, []*Node{&n1}}

	fmt.Println(cloneGraph(&n))
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	return copy(node, make(map[int]*Node))
}

func copy(node *Node, visited map[int]*Node) *Node {
	n, exists := visited[node.Val]
	if exists {
		return n
	}
	clone := Node{node.Val, nil}
	visited[node.Val] = &clone
	var neighbors []*Node
	for _, neighbor := range node.Neighbors {
		neighbors = append(neighbors, copy(neighbor, visited))

	}
	clone.Neighbors = neighbors
	return &clone
}
