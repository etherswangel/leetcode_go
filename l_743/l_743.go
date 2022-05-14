package leetcode

import "math"

func networkDelayTime(times [][]int, n int, k int) int {
	net := make([]map[int]int, n+1)
	for i := 0; i < len(net); i++ {
		net[i] = make(map[int]int, 0)
	}
	for _, time := range times {
		net[time[0]][time[1]] = time[2]
	}

	minPath, visited := make([]int, n+1), make([]bool, n+1)
	visited[k] = true

	for j := 0; j < n; j++ {
		for i, t := range visited {
			if !t {
				continue
			}
			for d, w := range net[i] {
				if visited[d] {
					if minPath[i]+w < minPath[d] {
						minPath[d] = minPath[i] + w
					}
				} else {
					visited[d] = true
					minPath[d] = minPath[i] + w
				}
			}
		}
	}

	for i := 1; i < len(visited); i++ {
		if !visited[i] {
			return -1
		}
	}

	res := 0
	for i := 1; i < len(minPath); i++ {
		if res < minPath[i] {
			res = minPath[i]
		}
	}
	return res
}

func networkDelayTime1(times [][]int, n int, k int) int {
	dist := make([]int, n+1)
	for i := 0; i < len(dist); i++ {
		dist[i] = math.MaxInt
	}
	dist[k] = 0
	for i := 0; i < n; i++ {
		for _, e := range times {
			u, v, w := e[0], e[1], e[2]
			if dist[u] != math.MaxInt && dist[v] > dist[u]+w {
				dist[v] = dist[u] + w
			}
		}
	}
	res := 0
	for i := 1; i < n+1; i++ {
		if dist[i] > res {
			res = dist[i]
		}
	}
	if res == math.MaxInt {
		return -1
	} else {
		return res
	}
}
