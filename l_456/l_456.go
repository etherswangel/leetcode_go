package leetcode

func find132pattern(nums []int) bool {
	stack := make([]int, 0)
	num_k := -2147483648

	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < num_k {
			return true
		}
		for len(stack) != 0 && stack[len(stack)-1] < nums[i] {
			num_k = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}
