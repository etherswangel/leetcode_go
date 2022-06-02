package leetcode

var int64Min, int64Max int = -1 << 63, 1<<63 - 1

func isValidBST(root *TreeNode) bool {
	return check(root, int64Min, int64Max)
}

func check(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	return min < root.Val && root.Val < max && check(root.Left, min, root.Val) && check(root.Right, root.Val, max)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
