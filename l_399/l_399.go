package leetcode

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	set := &stringSet{make(map[string]string, 0), make(map[string]float64, 0)}
	res := make([]float64, len(queries))
	for idx, eq := range equations {
		set.union(eq[0], eq[1], values[idx])
	}
	for idx, que := range queries {
		pX, vX, bX := set.find(que[0])
		pY, vY, bY := set.find(que[1])
		if bX == true || bY == true || pX != pY {
			res[idx] = -1
			continue
		}
		res[idx] = vX / vY
	}
	return res
}

type stringSet struct {
	parent map[string]string
	value  map[string]float64
}

func (s *stringSet) init(xs ...string) {
	for _, x := range xs {
		if _, ok := s.parent[x]; ok {
			continue
		}
		s.parent[x] = x
		s.value[x] = 1.0
	}
}

func (s *stringSet) find(x string) (parent string, value float64, err bool) {
	if _, ok := s.parent[x]; !ok {
		return x, -1, true
	}
	if x != s.parent[x] {
		px, vx, _ := s.find(s.parent[x])
		s.parent[x] = px
		s.value[x] *= vx
	}
	return s.parent[x], s.value[x], false
}

func (s *stringSet) union(a, b string, v float64) {
	s.init(a, b)

	pa, va, _ := s.find(a)
	pb, vb, _ := s.find(b)
	if pa != pb {
		s.parent[pa] = pb
		s.value[pa] = v * vb / va
	}
}
