package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	m, l, r := 0, 0, len(matrix[0])*len(matrix)-1

	for l <= r {
		m = l + (r-l)>>1
		if t := matrix[m/len(matrix[0])][m%len(matrix[0])]; t == target {
			return true
		} else if t < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return false
}
