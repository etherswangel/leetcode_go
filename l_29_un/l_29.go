package leetcode

const IntMin int = -2147483648
const IntMax int = 2147483647

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func divide(divided int, divisor int) int {
	if divided == IntMin && divisor == -1 {
		return IntMax
	}
	result := 0
	sign := -1
	if divided > 0 && divisor > 0 || divided < 0 && divisor < 0 {
		sign = 1
	}
	dvd, dvs := abs(divided), abs(divisor)
	for dvd >= dvs {
		temp := dvs
		m := 1
		for temp<<1 <= dvd {
			temp <<= 1
			m <<= 1
		}
		dvd -= temp
		result += m
	}
	return sign * result
}
