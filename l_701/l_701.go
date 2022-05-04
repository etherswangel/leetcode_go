package leetcode

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if root.Val == val {
			return root
		} else if root.Val < val {
			if root.Right != nil {
				insertIntoBST(root.Right, val)
			} else {
				root.Right = &TreeNode{val, nil, nil}
			}
		} else {
			if root.Left != nil {
				insertIntoBST(root.Left, val)
			} else {
				root.Left = &TreeNode{val, nil, nil}
			}
		}
	} else {
		root = &TreeNode{val, nil, nil}
	}
	return root
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
