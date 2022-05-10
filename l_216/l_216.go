package leetcode

func combinationSum3(k int, n int) [][]int {
	res := make([][]int, 0)
	curr := make([]int, 0)
	dfs(curr, 0, k, n, &res)
	return res
}

func dfs(curr []int, start, k, n int, res *[][]int) {
	if k == 0 && n == 0 {
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}
	for i := start + 1; i < 10 && i <= n; i++ {
		curr = append(curr, i)
		dfs(curr, i, k-1, n-i, res)
		curr = curr[:len(curr)-1]
	}
}
