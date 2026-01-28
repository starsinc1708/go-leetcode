package p117_populatingnextrightpointers

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	q := []*Node{root}

	for len(q) > 0 {
		level_size := len(q)

		for i := 0; i < level_size-1; i++ {
			q[i].Next = q[i+1]
		}

		for range level_size {
			node := q[0]
			q = q[1:]

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return root
}
