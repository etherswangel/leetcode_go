package leetcode

func backspaceCompare(s string, t string) bool {
	sS, tS := make([]rune, 0), make([]rune, 0)
	for _, c := range s {
		if c == '#' {
			if len(sS) > 0 {
				sS = sS[:len(sS)-1]
			}
		} else {
			sS = append(sS, c)
		}
	}
	for _, c := range t {
		if c == '#' {
			if len(tS) > 0 {
				tS = tS[:len(tS)-1]
			}
		} else {
			tS = append(tS, c)
		}
	}
	return string(sS) == string(tS)
}
