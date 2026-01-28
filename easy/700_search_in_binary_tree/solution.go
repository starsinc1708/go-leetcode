package p700_search_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root != nil && root.Val == val {
		return root
	}

	if root == nil {
		return nil
	}

	if val < root.Val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)
}
