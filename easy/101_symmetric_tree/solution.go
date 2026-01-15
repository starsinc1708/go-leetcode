package p101_symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root.Left)
	stack = append(stack, root.Right)

	for len(stack) > 0 {
		r := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		l := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if r == nil && l == nil {
			continue
		}
		if r == nil || l == nil {
			return false
		}
		if r.Val != l.Val {
			return false
		}

		stack = append(stack, r.Right)
		stack = append(stack, l.Left)
		stack = append(stack, r.Left)
		stack = append(stack, l.Right)
	}

	return true
}
