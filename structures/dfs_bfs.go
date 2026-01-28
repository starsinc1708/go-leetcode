package structures

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func dfs(root *TreeNode) {
	stack := root
	for stack != nil {
		node := stack
		stack = node.Right
		stack = node.Left
		print(node.Val)
	}
}

func bfs(root *TreeNode) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		level_size := len(q)
		for range level_size {
			node := q[0]
			q = q[1:]

			q = append(q, node.Left)
			q = append(q, node.Right)
		}
	}
}
