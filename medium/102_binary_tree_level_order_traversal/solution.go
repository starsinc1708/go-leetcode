package p102_binarytreelevelordertraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	result := make([][]int, 0)
	if root == nil {
		return result
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		level_size := len(q)
		curLevel := make([]int, 0)

		for range level_size {
			node := q[0]
			q = q[1:]
			curLevel = append(curLevel, node.Val)

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		result = append(result, curLevel)
	}

	return result
}
