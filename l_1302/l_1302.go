package leetcode

func deepestLeavesSum(root *TreeNode) int {
	s, q := 1, []*TreeNode{root}

	for len(q) != 0 {
		for i := 0; i < s; i++ {
			node := q[i]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if len(q) > s {
			q = q[s:]
			s = len(q)
		} else {
			res := 0
			for _, node := range q {
				res += node.Val
			}
			return res
		}
	}

	return -1
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
