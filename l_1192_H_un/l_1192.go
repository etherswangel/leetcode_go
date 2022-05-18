package leetcode

func criticalConnections(n int, connections [][]int) [][]int {
	disc, low := make([]int, n), make([]int, n)
	graph, res := make([][]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
		res[i] = make([]int, 0)
		disc[i] = -1
	}
	for i := 0; i < len(connections); i++ {
		from, to := connections[i][0], connections[i][1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}
	time := 0
	for i := 0; i < n; i++ {
		if disc[i] == -1 {
			dfs(i, low, disc, &graph, &res, i, &time)
		}
	}
	return res
}

func dfs(u int, low, disc []int, graph, res *[][]int, pre int, time *int) {
	*time++
	disc[u], low[u] = *time, *time
	for j := 0; j < len((*graph)[u]); j++ {
		v := (*graph)[u][j]
		if v == pre {
			continue
		}
		if disc[v] == -1 {
			dfs(v, low, disc, graph, res, u, time)
			low[u] = min(low[u], low[v])
			if low[v] > disc[u] {
				*res = append(*res, []int{u, v})
			}
		} else {
			low[u] = min(low[u], disc[v])
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
