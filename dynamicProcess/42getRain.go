package goleet

// https://leetcode-cn.com/problems/trapping-rain-water/

func trap(height []int) int {
	lenh := len(height)
	var dp []int = make([]int, lenh)
	dp[0] = 0
	for i := 1; i < lenh; i++ {
		if height[i] <= height[i-1] {
			dp[i] = dp[i-1]
		} else {
			maxh := height[i-1]
			dp[i] = dp[i-1]
			for j := i - 2; j >= 0; j-- {
				if height[j] <= maxh {
					continue
				} else if height[j] > maxh && height[j] < height[i] {
					dp[i] += (height[j] - maxh) * (i - j - 1)
					maxh = height[j]
				} else {
					dp[i] += (height[i] - maxh) * (i - j - 1)
					break
				}
			}
		}
	}
	return dp[lenh-1]
}
