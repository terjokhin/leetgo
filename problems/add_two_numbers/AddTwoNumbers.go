package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{Val: 9, Next: nil}
	l2 := ListNode{Val: 4, Next: &l1}
	l3 := ListNode{Val: 2, Next: &l2}

	r1 := ListNode{Val: 9, Next: nil}
	r2 := ListNode{Val: 4, Next: &r1}
	r3 := ListNode{Val: 6, Next: &r2}
	r4 := ListNode{Val: 5, Next: &r3}

	result := addTwoNumbers(&l3, &r4)
	for result != nil {
		log.Println(result.Val)
		result = result.Next
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0

	var buff []int

	for l1 != nil || l2 != nil {
		v1 := 0
		if l1 != nil {
			v1 = l1.Val
		}

		v2 := 0
		if l2 != nil {
			v2 = l2.Val
		}

		sum := v1 + v2 + carry

		carry = sum / 10
		v := sum % 10

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

		buff = append(buff, v)

	}

	if carry != 0 {
		buff = append(buff, carry)
	}

	var r *ListNode

	for i := len(buff); i > 0; i-- {
		r = &ListNode{Val: buff[i-1], Next: r}
	}

	return r
}
