package p1302_deepestleavessum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}

	result := 0

	for len(q) > 0 {
		level_size := len(q)
		result = 0
		for range level_size {

			node := q[0]
			q = q[1:]

			result += node.Val

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

	}

	return result
}
