package leetcode

import "sort"

func findUnsortedSubarray(nums []int) int {
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	sort.Ints(sortedNums)

	l, r := 0, len(nums)-1
	for l <= r && nums[l] == sortedNums[l] {
		l++
	}
	for l <= r && nums[r] == sortedNums[r] {
		r--
	}
	return r - l + 1
}

func findUnsortedSubarray1(nums []int) int {
	start, end, max := -1, -1, -1000000
	for i := 0; i < len(nums); i++ {
		if max > nums[i] {
			if start == -1 {
				start = i - 1
			}
			for start > 0 && nums[start-1] > nums[i] {
				start--
			}
			end = i + 1
		} else {
			max = nums[i]
		}
	}
	return end - start
}
