package leetcode

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	size, dp := 0, make([]int, len(envelopes))
	for _, envelope := range envelopes {
		l, r, m := 0, size, 0
		for l < r {
			m = l + (r-l)>>1
			if dp[m] < envelope[1] {
				l = m + 1
			} else {
				r = m
			}
		}
		dp[l] = envelope[1]
		if l == size {
			size++
		}
	}
	return size
}
