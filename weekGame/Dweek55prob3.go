package weekgame

import "math"

// https://leetcode-cn.com/contest/biweekly-contest-55/problems/maximum-alternating-subsequence-sum/

func maxAlternatingSum(nums []int) int64 {
	lenN := len(nums)
	type pairN struct {
		cur int
		low int
		sum int
	}
	dp := make([]pairN, lenN)
	dp[0] = pairN{nums[0], math.MaxInt32, nums[0]}
	for i := 1; i < lenN; i++ {
		n := nums[i]
		cur, low, sum := dp[i-1].cur, dp[i-1].low, dp[i-1].sum
		if n > low {
			dp[i].cur = n
			dp[i].low = math.MaxInt32
			dp[i].sum = sum + n - low
		} else {
			if n >= cur {
				dp[i].cur = n
				dp[i].low = low
				dp[i].sum = sum + n - cur
			} else {
				dp[i].low = n
				dp[i].sum = sum
				dp[i].cur = cur
			}
		}
	}
	return int64(dp[lenN-1].sum)
}
