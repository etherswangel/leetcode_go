package leetcode

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func groupAnagrams(strs []string) [][]string {
	m, res := make(map[string][]string), make([][]string, 0)

	for _, s := range strs {
		r := sortRunes(s)
		sort.Sort(r)
		ss := m[string(r)]
		ss = append(ss, s)
		m[string(r)] = ss
	}

	for _, v := range m {
		res = append(res, v)
	}

	return res
}
