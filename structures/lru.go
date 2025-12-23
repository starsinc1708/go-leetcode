package structures

type Node struct {
	Key   int
	Value int
	Next  *Node
	Prev  *Node
}

type LRUCache struct {
	Capacity int
	Data     map[int]*Node
	Head     *Node
	Tail     *Node
}

func LRUCacheConstructor(capacity int) LRUCache {
	data := make(map[int]*Node, capacity)

	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		Capacity: capacity,
		Data:     data,
		Head:     head,
		Tail:     tail,
	}
}

func (this *LRUCache) remove(node *Node) {
	if node == this.Head || node == this.Tail {
		return
	}

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) AddToHead(node *Node) {
	node.Next = this.Head.Next
	this.Head.Next.Prev = node

	node.Prev = this.Head
	this.Head.Next = node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Data[key]
	if !ok {
		return -1
	}
	this.remove(node)
	this.AddToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Data[key]
	if ok {
		this.remove(node)
		this.AddToHead(node)
		node.Value = value
		return
	}

	if len(this.Data) >= this.Capacity {
		node := this.Tail.Prev
		this.remove(node)
		delete(this.Data, node.Key)
	}

	node = &Node{Key: key, Value: value}

	this.AddToHead(node)
	this.Data[key] = node
}
