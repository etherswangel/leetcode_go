package leetcode

func longestIncreasingPath(matrix [][]int) int {
	m, n, res := len(matrix), len(matrix[0]), 0
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}

	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			res = max(res, bfs(&matrix, &dirs, &visited, m, n, x, y))
		}
	}

	return res
}

func bfs(matrix, dirs, visited *[][]int, m, n, x, y int) int {
	if (*visited)[x][y] != 0 {
		return (*visited)[x][y]
	}

	res := 0
	for _, dir := range *dirs {
		_x, _y := x+dir[0], y+dir[1]
		if valid(m, n, _x, _y) && (*matrix)[x][y] < (*matrix)[_x][_y] {
			res = max(res, bfs(matrix, dirs, visited, m, n, _x, _y))
		}
	}
	(*visited)[x][y] = res + 1
	return res + 1
}

func valid(m, n, x, y int) bool {
	if x < 0 || y < 0 || x >= m || y >= n {
		return false
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
