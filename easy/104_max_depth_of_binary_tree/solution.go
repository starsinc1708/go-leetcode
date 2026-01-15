package p104_maxdepthofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := maxDepth(root.Left)
	r := maxDepth(root.Right)

	return max(l, r) + 1
}

type NodeDepth struct {
	Node  *TreeNode
	Depth int
}

func maxDepthWithoutRecursion(root *TreeNode) int {
	stack := []NodeDepth{}
	stack = append(stack, NodeDepth{Node: root, Depth: 1})
	result := 0

	for len(stack) > 0 {
		node := stack[len(stack)-1].Node
		depth := stack[len(stack)-1].Depth

		stack = stack[:len(stack)-1]

		if node == nil {
			continue
		}

		result = max(result, depth)
		stack = append(stack, NodeDepth{node.Left, depth + 1})
		stack = append(stack, NodeDepth{node.Right, depth + 1})
	}

	return result
}
