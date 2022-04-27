package leetcode

import "sort"

type Set struct {
	parent []int
}

func (set *Set) Init(n int) {
	set.parent = make([]int, n)
	for i := 0; i < n; i++ {
		set.parent[i] = i
	}
}

func (set *Set) Find(x int) int {
	if x != set.parent[x] {
		set.parent[x] = set.Find(set.parent[x])
	}
	return set.parent[x]
}

func (set *Set) Union(x, y int) {
	x = set.Find(x)
	y = set.Find(y)
	if x == y {
		return
	}
	set.parent[x] = y
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	if len(s) == 1 || len(pairs) == 0 {
		return s
	}

	set := &Set{}
	set.Init(len(s))

	sMap := make(map[int][]byte, 0)

	for _, pair := range pairs {
		set.Union(pair[0], pair[1])
	}

	for idx := range s {
		root := set.Find(idx)
		sMap[root] = append(sMap[root], s[idx])
	}

	for _, s := range sMap {
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
	}

	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		root := set.Find(i)
		res[i] = sMap[root][0]
		sMap[root] = sMap[root][1:]
	}

	return string(res)
}
