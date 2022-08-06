package greedy

import "sort"

// https://leetcode.cn/problems/minimum-subsequence-in-non-increasing-order/

func minSubsequence(nums []int) (ans []int) {
	sort.Ints(nums)
	sum, tmp := 0, 0
	for _, v := range nums {
		sum += v
	}
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		tmp += nums[i]
		ans = append(ans, nums[i])
		if tmp > sum/2 {
			break
		}
	}
	return
}
