package leetcode

import "bytes"

type pair struct {
	char  byte
	count int
}

func removeDuplicates(s string, k int) string {
	stack := make([]pair, 0)
	for i := range s {
		if len(stack) == 0 || stack[len(stack)-1].char != s[i] {
			stack = append(stack, pair{s[i], 1})
		} else {
			stack[len(stack)-1].count++
		}
		if stack[len(stack)-1].count == k {
			stack = stack[:len(stack)-1]
		}
	}
	bb := new(bytes.Buffer)
	for _, c := range stack {
		for i := 0; i < c.count; i++ {
			bb.WriteByte(c.char)
		}
	}
	return bb.String()
}
