package mathgame

import "sort"

// https://leetcode.cn/problems/smallest-range-i/

func smallestRangeI(nums []int, k int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	sort.Ints(nums)
	dif := nums[n-1] - nums[0]
	if 2*k >= dif {
		return 0
	} else {
		return dif - 2*k
	}
}
