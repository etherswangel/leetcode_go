package leetcode

type union struct {
	pa []int
	ra []int
}

func equationsPossible(equations []string) bool {
	set := union{}
	set.init(26)

	for _, s := range equations {
		if s[1] == '=' {
			set.merge(int(s[0]-'a'), int(s[3]-'a'))
		}
	}

	for _, s := range equations {
		if s[1] == '!' {
			if set.find(int(s[0]-'a')) == set.find(int(s[3]-'a')) {
				return false
			}
		}
	}

	return true
}

func (p *union) init(n int) {
	p.pa, p.ra = make([]int, n), make([]int, n)

	for i := 0; i < n; i++ {
		p.pa[i] = i
		p.ra[i] = 1
	}
}

func (p *union) find(a int) int {
	if a == p.pa[a] {
		return a
	} else {
		p.pa[a] = p.find(p.pa[a])
		return p.pa[a]
	}
}

func (p *union) merge(a, b int) {
	x, y := p.find(a), p.find(b)

	if x == y {
		return
	}

	if p.ra[x] < p.ra[y] {
		p.pa[x] = y
	} else if p.ra[x] > p.ra[y] {
		p.pa[y] = x
	} else {
		p.pa[x] = y
		p.ra[y] += 1
	}
}
