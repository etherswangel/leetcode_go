package leetcode

const intMax = int(^uint(0) >> 1)

func minCostConnectPoints(points [][]int) int {
	res := 0
	distances := make([]int, len(points))
	currIndex := 0

	for currIndex >= 0 {
		distances[currIndex] = -1
		minEdge := 0
		minVertex := -1
		for i := 0; i < len(points); i++ {
			if distances[i] == -1 {
				continue
			}
			d := distance(points[currIndex], points[i])
			if distances[i] == 0 || d < distances[i] {
				distances[i] = d
			}
			if minEdge == 0 || distances[i] < minEdge {
				minEdge = distances[i]
				minVertex = i
			}
		}
		res += minEdge
		currIndex = minVertex
	}
	return res
}

func distance(a, b []int) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
