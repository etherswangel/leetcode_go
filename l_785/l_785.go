package leetcode

func isBipartite(graph [][]int) bool {
	size := len(graph)
	set := &Set{}
	set.init(size)

	for id, it := range graph {
		if len(it) == 0 {
			continue
		}
		for _, n := range it {
			if set.isConnected(id, n) {
				return false
			}
			set.union(it[0], n)
		}
	}

	return true
}

type Set struct {
	parent []int
	count  int
}

func (s *Set) init(n int) {
	s.count = n
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

func (s *Set) union(a, b int) {
	a, b = s.find(a), s.find(b)
	if a != b {
		s.parent[a] = b
		s.count--
	}
}

func (s *Set) isConnected(a, b int) bool {
	a, b = s.find(a), s.find(b)
	if a == b {
		return true
	}
	return false
}
