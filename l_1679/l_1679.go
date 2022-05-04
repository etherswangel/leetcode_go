package leetcode

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	i, j, res := 0, len(nums)-1, 0
	for i < j {
		if sum := nums[i] + nums[j]; sum == k {
			res++
			i++
			j--
		} else if sum < k {
			i++
		} else {
			j--
		}
	}
	return res
}

func maxOperations1(nums []int, k int) int {
	freq := make(map[int]int, 0)
	res := 0
	for _, n := range nums {
		if del := k - n; freq[del] > 0 {
			res++
			freq[del]--
		} else {
			freq[n]++
		}
	}
	return res
}
