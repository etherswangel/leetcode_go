package leetcode

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 || grid[1][1] == 1 {
		return -1
	}

	dirs := [][]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	q, n := [][]int{{0, 0, 1}}, len(grid)

	for len(q) != 0 {
		x, y, d := q[0][0], q[0][1], q[0][2]+1
		q = q[1:]
		if x == n-1 && y == n-1 {
			return d
		}
		for _, dir := range dirs {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 0 {
				grid[nx][ny] = 1
				q = append(q, []int{nx, ny, d})
			}
		}
	}
	return -1
}
