package p141_linked_list_cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// Алгоритм Флойда
func hasCycle(head *ListNode) bool {
	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
