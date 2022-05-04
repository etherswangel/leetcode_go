package leetcode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if root.Val < val {
			return searchBST(root.Right, val)
		} else if root.Val > val {
			return searchBST(root.Left, val)
		}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
