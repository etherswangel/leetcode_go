package leetcode

import "fmt"

func hasAllCodes(s string, k int) bool {
	if len(s) < k {
		return false
	}

	m_map := make(map[string]bool, 0)
	for i := 0; i < len(s)-k+1; i++ {
		m_map[s[i:i+k]] = true
		fmt.Println(i, i+k)
	}
	fmt.Println(m_map)

	return len(m_map) == (1 << k)
}
