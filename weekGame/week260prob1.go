package weekgame

import "math"

// https://leetcode-cn.com/contest/weekly-contest-260/problems/maximum-difference-between-increasing-elements/

func maximumDifference(nums []int) int {
	ans, lenN := -1, len(nums)
	if lenN <= 1 {
		return ans
	}
	num1 := max(nums[1:]) - nums[0]
	if num1 > 0 {
		ans = num1
	}
	tmp := math.MaxInt32
	for i := 1; i < lenN-1; i++ {
		if nums[i] < tmp && nums[i] < nums[i+1] && nums[i] < nums[i-1] {
			numi := max(nums[i+1:]) - nums[i]
			if ans < numi {
				ans = numi
			}
			tmp = nums[i]
		}
	}
	return ans
}

func max(nums []int) int {
	res := math.MinInt32
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}
