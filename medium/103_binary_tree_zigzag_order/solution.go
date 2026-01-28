package p103_binarytreezigzagorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	if root == nil {
		return result
	}

	q := []*TreeNode{root}
	flag := true

	for len(q) > 0 {
		level_size := len(q)
		curLevel := make([]int, 0)

		for range level_size {
			if flag {
				node := q[0]
				q = q[1:]

				curLevel = append(curLevel, node.Val)

				if node.Left != nil {
					q = append(q, node.Left)
				}
				if node.Right != nil {
					q = append(q, node.Right)
				}

			} else {

				node := q[len(q)-1]
				q = q[:len(q)-1]

				curLevel = append(curLevel, node.Val)

				if node.Right != nil {
					q = append([]*TreeNode{node.Right}, q...)
				}
				if node.Left != nil {
					q = append([]*TreeNode{node.Left}, q...)
				}
			}
		}

		flag = !flag
		result = append(result, curLevel)
	}

	return result
}
