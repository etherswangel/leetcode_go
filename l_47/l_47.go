package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	n, res := len(nums), make([][]int, 0)

	t := make([]int, n)
	copy(t, nums)
	res = append(res, t)
	for {
		j, k := n-1, n-1
		for j > 0 && nums[j-1] >= nums[j] {
			j--
		}
		if j == 0 {
			break
		}
		for k >= j && nums[k] <= nums[j-1] {
			k--
		}
		nums[j-1], nums[k] = nums[k], nums[j-1]

		tmp, cur := make([]int, n-j), make([]int, n)
		copy(tmp, nums[j:n])
		sort.Ints(tmp)
		copy(nums[j:n], tmp)
		copy(cur, nums)

		res = append(res, cur)
	}

	return res
}
