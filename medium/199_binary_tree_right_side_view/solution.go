package p_199_binarytreerightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)

	if root == nil {
		return result
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		level_size := len(q)

		result = append(result, q[len(q)-1].Val)

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

	return result
}
