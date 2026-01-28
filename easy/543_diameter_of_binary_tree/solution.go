package p543_diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var d int

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	d = max(d, l+r)

	return max(l, r) + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	d = 0

	maxDepth(root)
	return d
}

func diameterOfBinaryTreeAlt(root *TreeNode) int {
	var dfs func(*TreeNode) (depth, d int)

	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}

		l, ld := dfs(root.Left)
		r, rd := dfs(root.Right)

		depth := max(l, r) + 1
		d := max(max(ld, rd), l+r)

		return depth, d
	}

	_, diameter := dfs(root)

	return diameter
}
