package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	node1 := ListNode{1, nil}
	node2 := ListNode{2, &node1}
	node3 := ListNode{3, &node2}
	node4 := ListNode{4, &node3}

	traverse(sortList(&node4))
}

func traverse(l *ListNode) {
	curr := l

	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}

}

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	middle := getMiddle(head)
	next := middle.Next

	middle.Next = nil

	left := sortList(head)
	right := sortList(next)

	return sortAndMerge(left, right)
}

func sortAndMerge(l *ListNode, r *ListNode) *ListNode {

	if l == nil {
		return r
	}

	if r == nil {
		return l
	}

	var result ListNode

	if l.Val < r.Val {
		result.Val = l.Val
		result.Next = sortAndMerge(l.Next, r)
	} else {
		result.Val = r.Val
		result.Next = sortAndMerge(l, r.Next)
	}

	return &result
}

func getMiddle(node *ListNode) *ListNode {
	if node == nil {
		return node
	}

	slow, fast := node, node

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
