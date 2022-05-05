package leetcode

func findTarget(root *TreeNode, k int) bool {
	lStack, rStack := make([]*TreeNode, 0), make([]*TreeNode, 0)
	left, right := root, root
	lStack = append(lStack, left)
	rStack = append(rStack, right)

	for left.Left != nil {
		left = left.Left
		lStack = append(lStack, left)
	}
	for right.Right != nil {
		right = right.Right
		rStack = append(rStack, right)
	}

	for left != right {
		if sum := left.Val + right.Val; sum == k {
			return true
		} else if sum < k {
			left = next(&lStack, left)
		} else {
			right = prev(&rStack, right)
		}
	}
	return false
}

func next(stack *[]*TreeNode, root *TreeNode) *TreeNode {
	*stack = (*stack)[:len(*stack)-1]
	if root.Right != nil {
		root = root.Right
		*stack = append(*stack, root)
		for root.Left != nil {
			root = root.Left
			*stack = append(*stack, root)
		}
	} else {
		root = (*stack)[len(*stack)-1]
	}
	return root
}

func prev(stack *[]*TreeNode, root *TreeNode) *TreeNode {
	*stack = (*stack)[:len(*stack)-1]
	if root.Left != nil {
		root = root.Left
		*stack = append(*stack, root)
		for root.Right != nil {
			root = root.Right
			*stack = append(*stack, root)
		}
	} else {
		root = (*stack)[len(*stack)-1]
	}
	return root
}

//////////////////////////
func findTarget1(root *TreeNode, k int) bool {
	m := make(map[int]int, 0)
	return findTargetDFS(root, k, m)
}

func findTargetDFS(root *TreeNode, k int, m map[int]int) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return ok
	}
	m[root.Val]++
	return findTargetDFS(root.Left, k, m) || findTargetDFS(root.Right, k, m)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
