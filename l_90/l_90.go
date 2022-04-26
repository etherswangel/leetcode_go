package leetcode

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 1)
	tmp := make([]int, 0)

	generateSubsetsWithDup(&nums, 0, tmp, &res)

	return res
}

func generateSubsetsWithDup(nums *[]int, start int, curr []int, res *[][]int) {
	for i := start; i < len(*nums); i++ {
		if i != start && (*nums)[i] == (*nums)[i-1] {
			continue
		}
		curr = append(curr, (*nums)[i])
		*res = append(*res, append([]int{}, curr...))
		generateSubsetsWithDup(nums, i+1, curr, res)
		curr = curr[:len(curr)-1]
	}
}
