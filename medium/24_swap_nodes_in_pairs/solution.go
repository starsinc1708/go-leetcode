package p24_swap_nodes_in_pairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		second := cur.Next.Next

		first.Next = second.Next
		second.Next = first
		cur.Next = second
		cur = first
	}

	return dummy.Next
}
