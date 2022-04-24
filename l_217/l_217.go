package leetcode

import "sort"

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for _, n := range nums {
		if m[n] {
			return true
		}
		m[n] = true
	}
	return false
}

func containsDuplicate2(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}
