package leetcode

func removeDuplicates(nums []int) int {
	n := 0
	for _, i := range nums {
		if n < 2 || i > nums[n-2] {
			nums[n] = i
			n++
		}
	}
	return n
}
