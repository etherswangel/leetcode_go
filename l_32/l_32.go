package leetcode

func longestValidParentheses(s string) int {
	stack := make([]pair, 0)
	start, end, max := 0, 0, 0
	for idx := range s {
		ch := s[idx]
		if ch == ')' {
			if len(stack) == 0 {
				continue
			}
			if stack[len(stack)-1].ch == '(' {
				end = idx
				stack = stack[:len(stack)-1]
			} else {
				end = idx
				if del := end - start; max < del {
					max = del
				}
				stack = make([]pair, 0)
				start, end = idx+1, idx+1
			}
		} else {
			stack = append(stack, pair{s[idx], idx})
		}
	}
	if len(stack) != 0 {
		start = stack[len(stack)-1].idx + 1
	}
	if del := len(s) - start; max < del {
		max = del
	}
	return max
}

type pair struct {
	ch  byte
	idx int
}
