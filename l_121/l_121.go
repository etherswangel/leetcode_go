package leetcode

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if del := prices[i] - min; del > max {
			max = del
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return max
}
