package arraymanipulate

import "math"

// https://leetcode.cn/problems/third-maximum-number/submissions/

func thirdMax(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	tmp := [3]int{math.MinInt64, math.MinInt64, math.MinInt64}
	for _, v := range nums {
		if v == tmp[2] || v == tmp[1] {
			continue
		}
		if v > tmp[2] {
			tmp[0] = tmp[1]
			tmp[1] = tmp[2]
			tmp[2] = v
		} else if v > tmp[1] {
			tmp[0] = tmp[1]
			tmp[1] = v
		} else if v > tmp[0] {
			tmp[0] = v
		}
	}
	if tmp[0] != math.MinInt64 {
		return tmp[0]
	} else {
		return tmp[2]
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
