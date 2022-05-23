package leetcode

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < len(strs); i++ {
		zs, os := 0, 0
		for _, c := range strs[i] {
			if c == '0' {
				zs++
			}
			if c == '1' {
				os++
			}
		}

		for j := m; j >= zs; j-- {
			for k := n; k >= os; k-- {
				if dp[j][k] < dp[j-zs][k-os]+1 {
					dp[j][k] = dp[j-zs][k-os] + 1
				}
			}
		}
	}
	return dp[m][n]
}
