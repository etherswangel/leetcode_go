package leetcode

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return make([]string, 0)
	}

	dict := [][]string{{}, {}, {"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}

	res := make([]string, 0)
	res = append(res, "")
	for idx := range digits {
		num := digits[idx] - '0'
		length := len(res)
		for i := 0; i < length; i++ {
			for _, c := range dict[num] {
				res = append(res, res[i]+c)
			}
		}
		res = res[length:]
	}
	return res
}
