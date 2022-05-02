package leetcode

func sortArrayByParity(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		for i < j && nums[i]&1 != 1 {
			i++
		}
		for i < j && nums[j]&1 == 1 {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
