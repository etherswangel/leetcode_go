package leetcode

func isSymmetric(root *TreeNode) bool {
	leftQueue, rightQueue := make([]*TreeNode, 0), make([]*TreeNode, 0)
	leftQueue = append(leftQueue, root)
	rightQueue = append(rightQueue, root)

	for len(leftQueue) > 0 && len(leftQueue) == len(rightQueue) {
		left, right := leftQueue[0], rightQueue[0]
		leftQueue = leftQueue[1:]
		rightQueue = rightQueue[1:]

		if left == nil && right == nil {
			continue
		} else if left != nil && right != nil {
			if left.Val == right.Val {
				leftQueue = append(leftQueue, left.Left)
				leftQueue = append(leftQueue, left.Right)
				rightQueue = append(rightQueue, right.Right)
				rightQueue = append(rightQueue, right.Left)
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
