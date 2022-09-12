package arraymanipulate

import "sort"

// https://leetcode.cn/problems/special-array-with-x-elements-greater-than-or-equal-x/

func specialArray(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	if nums[l-1] < 0 {
		return 0
	}
	for i := 0; i < l; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if i > 0 && nums[i] >= l-i && nums[i-1] < l-i {
			return l - i
		}
		if i == 0 && nums[i] >= l {
			return l
		}
	}
	return -1
}
