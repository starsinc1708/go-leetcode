package p112_pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeTarget struct {
	Node *TreeNode
	Sum  int
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := make([]*NodeTarget, 0)

	stack = append(stack, &NodeTarget{Node: root, Sum: 0})

	for len(stack) > 0 {
		node := stack[len(stack)-1].Node
		curSum := stack[len(stack)-1].Sum
		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		curSum += node.Val

		if node.Left == nil && node.Right == nil && curSum == targetSum {
			return true
		}

		stack = append(stack, &NodeTarget{
			Node: node.Left,
			Sum:  curSum,
		})

		stack = append(stack, &NodeTarget{
			Node: node.Right,
			Sum:  curSum,
		})
	}
	return false
}
