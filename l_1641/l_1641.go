package leetcode

func countVowelStrings(n int) int {
	dic := []int{1, 1, 1, 1, 1}
	for n > 1 {
		dic[4] += dic[0] + dic[1] + dic[2] + dic[3]
		dic[3] += dic[0] + dic[1] + dic[2]
		dic[2] += dic[0] + dic[1]
		dic[1] += dic[0]
		n--
	}
	return dic[0] + dic[1] + dic[2] + dic[3] + dic[4]
}
