package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	input4 := ListNode{4, nil}
	input3 := ListNode{3, &input4}
	input2 := ListNode{2, &input3}
	input1 := ListNode{1, &input2}
	fmt.Println(swapPairs(&input1))
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := swapPairs(head.Next.Next)
	temp := head.Next
	head.Next = nextHead
	temp.Next = head
	return temp
}
