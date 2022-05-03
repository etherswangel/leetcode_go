package leetcode

func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertTree(root.Left)
		invertTree(root.Right)
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
