package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	m := make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}

	res := make([]int, 0)
	for _, num := range nums2 {
		if m[num] > 0 {
			res = append(res, num)
			m[num]--
		}
	}
	return res
}
