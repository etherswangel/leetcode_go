package leetcode

func levelOrder(root *TreeNode) [][]int {
	res, pos := make([][]int, 0, 2000), make([]*TreeNode, 0, 2000)
	pos = append(pos, root)

	for len(pos) > 0 {
		start := len(pos)
		tmp := make([]int, 0, start)

		for i := 0; i < start; i++ {
			if pos[0] != nil {
				tmp = append(tmp, pos[0].Val)
				pos = append(pos, pos[0].Left)
				pos = append(pos, pos[0].Right)
			}
			pos = pos[1:]
		}
		res = append(res, tmp)
	}
	return res[:len(res)-1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
