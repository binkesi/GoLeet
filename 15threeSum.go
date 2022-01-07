package goleet

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	numLen := len(nums)
	if numLen < 3 {
		return res
	}
	for i := 0; i < numLen; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		twoSum := 0 - nums[i]
		for j, k := i+1, numLen-1; j < numLen; j++ {
			if j >= k {
				break
			}
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > twoSum {
				k--
			}
			if j >= k {
				continue
			} else {
				if nums[j]+nums[k] == twoSum {
					tmp := []int{}
					tmp = append(tmp, nums[i], nums[j], nums[k])
					res = append(res, tmp)
					k--
				}
				continue
			}
		}
	}
	return res
}
