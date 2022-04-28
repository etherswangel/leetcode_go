package leetcode

import "sort"

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	edges := make([]edge, 0)
	for i, row := range heights {
		for j := range row {
			id := i*n + j
			if i > 0 {
				edges = append(edges, edge{id - n, id, abs(heights[i][j] - heights[i-1][j])})
			}
			if j > 0 {
				edges = append(edges, edge{id - 1, id, abs(heights[i][j] - heights[i][j-1])})
			}
		}
	}
	sort.Slice(edges, func(a, b int) bool {
		return edges[a].w < edges[b].w
	})

	set := &Set{}
	set.init(m * n)
	for _, it := range edges {
		set.union(it.v1, it.v2)
		if set.find(0) == set.find(m*n-1) {
			return it.w
		}
	}
	return 0
}

type edge struct {
	v1, v2, w int
}

type Set struct {
	parent []int
}

func (s *Set) init(n int) {
	s.parent = make([]int, n)
	for i := range s.parent {
		s.parent[i] = i
	}
}

func (s *Set) find(x int) int {
	if s.parent[x] != x {
		s.parent[x] = s.find(s.parent[x])
	}
	return s.parent[x]
}

func (s *Set) union(x, y int) {
	pX, pY := s.find(x), s.find(y)
	if pX != pY {
		s.parent[pX] = pY
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/**
 * DFS+BS
**/
var dir = [4][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}
	low, high := 0, 1000000
	for low < high {
		threshold := low + (high-low)>>1
		if !hasPath(heights, visited, 0, 0, threshold) {
			low = threshold + 1
		} else {
			high = threshold
		}
		for i := range visited {
			for j := range visited[i] {
				visited[i][j] = false
			}
		}
	}
	return low
}

func hasPath(heights [][]int, visited [][]bool, i, j, threshold int) bool {
	n, m := len(heights), len(heights[0])
	if i == n-1 && j == m-1 {
		return true
	}
	visited[i][j] = true
	for _, d := range dir {
		ni, nj := i+d[0], j+d[1]
		if ni < 0 || ni >= n || nj < 0 || nj >= m || visited[ni][nj] {
			continue
		}
		diff := abs(heights[i][j] - heights[ni][nj])
		if diff <= threshold && hasPath(heights, visited, ni, nj, threshold) {
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
