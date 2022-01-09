package goleet

// https://leetcode-cn.com/problems/longest-increasing-subsequence/

func lengthOfLIS(nums []int) int {
	lenn := len(nums)
	var dp []int = make([]int, lenn)
	dp[0] = 1
	for i := 1; i < lenn; i++ {
		dp[i] = 1
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] {
				dp[i] = maxN(dp[i], dp[j]+1)
			}
		}
	}
	maxl := 0
	for _, v := range dp {
		if v > maxl {
			maxl = v
		}
	}
	return maxl
}

func maxN(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
