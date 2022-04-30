package leetcode

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if ch := s[i]; ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if len(stack) > 0 && ch-stack[len(stack)-1] <= 2 {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}
