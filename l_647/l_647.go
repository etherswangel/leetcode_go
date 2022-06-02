package leetcode

import "fmt"

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		j := 0
		for j < len(s)-i && s[i+j] == s[i] {
			j++
		}
		res += (j * (j + 1)) / 2
		fmt.Println("a", res)

		k := i + j - 1
		for l := 1; ; l++ {
			if i-l >= 0 && k+l < len(s) && s[i-l] == s[k+l] {
				res++
			} else {
				break
			}
		}
		i += j - 1
		fmt.Println(res)
	}
	return res
}
