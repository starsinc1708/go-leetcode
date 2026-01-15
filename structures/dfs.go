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
	return
}
