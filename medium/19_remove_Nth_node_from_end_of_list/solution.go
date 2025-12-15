package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}

	fast, slow := dummy, dummy

	for i := n + 1; i > 0; i-- {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}

func main() {
	fmt.Println(removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 2))
}
