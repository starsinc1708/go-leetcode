package structures

type ListNode struct {
	Value int
	Next  *ListNode
}

type MyLinkedList struct {
	Head *ListNode
	Size int
}

func LinkedListConstructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	current := this.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}

	this.Size++

	// вставка в голову
	if index == 0 {
		newNode := &ListNode{Value: val, Next: this.Head}
		this.Head = newNode
		return
	}

	current := this.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}

	newNode := &ListNode{Value: val, Next: current.Next}
	current.Next = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}

	this.Size--

	if index == 0 {
		this.Head = this.Head.Next
		return
	}

	current := this.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
}
