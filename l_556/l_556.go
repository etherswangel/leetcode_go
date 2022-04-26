package leetcode

func matrixReshape(mat [][]int, r int, c int) [][]int {
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}
	idx, res := 0, make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			res[i][j] = mat[idx/len(mat[0])][idx%len(mat[0])]
			idx++
		}
	}
	return res
}
