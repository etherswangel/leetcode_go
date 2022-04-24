package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for idx_1 := 0; idx_1 < len(nums)-3; idx_1++ {
		if idx_1 > 0 && nums[idx_1] == nums[idx_1-1] {
			continue
		}

		for idx_2 := idx_1 + 1; idx_2 < len(nums)-2; idx_2++ {
			if idx_2 > idx_1+1 && nums[idx_2] == nums[idx_2-1] {
				continue
			}

			tgt := target - nums[idx_1] - nums[idx_2]
			for idx_3, idx_4 := idx_2+1, len(nums)-1; idx_3 < idx_4; {
				if sum := nums[idx_3] + nums[idx_4]; sum == tgt {
					res = append(res, []int{nums[idx_1], nums[idx_2], nums[idx_3], nums[idx_4]})
					idx_3++
					idx_4--
					for idx_3 < idx_4 && nums[idx_3] == nums[idx_3-1] {
						idx_3++
					}
					for idx_3 < idx_4 && nums[idx_4] == nums[idx_4+1] {
						idx_4--
					}
				} else if sum < tgt {
					idx_3++
				} else {
					idx_4--
				}
			}
		}
	}
	return res
}
