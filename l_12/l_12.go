package leetcode

import "strings"

func intToRoman(num int) string {
	var ans strings.Builder
	ans.Grow(10)

	for num > 0 {
		if num >= 1000 {
			ans.WriteByte('M')
			num -= 1000
		} else if num >= 900 {
			ans.WriteByte('C')
			ans.WriteByte('M')
			num -= 900
		} else if num >= 500 {
			ans.WriteByte('D')
			num -= 500
		} else if num >= 400 {
			ans.WriteByte('C')
			ans.WriteByte('D')
			num -= 400
		} else if num >= 100 {
			ans.WriteByte('C')
			num -= 100
		} else if num >= 90 {
			ans.WriteByte('X')
			ans.WriteByte('C')
			num -= 90
		} else if num >= 50 {
			ans.WriteByte('L')
			num -= 50
		} else if num >= 40 {
			ans.WriteByte('X')
			ans.WriteByte('L')
			num -= 40
		} else if num >= 10 {
			ans.WriteByte('X')
			num -= 10
		} else if num >= 9 {
			ans.WriteByte('I')
			ans.WriteByte('X')
			num -= 9
		} else if num >= 5 {
			ans.WriteByte('V')
			num -= 5
		} else if num >= 4 {
			ans.WriteByte('I')
			ans.WriteByte('V')
			num -= 4
		} else {
			ans.WriteByte('I')
			num--
		}
	}

	return ans.String()
}
