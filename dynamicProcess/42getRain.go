package dynamicprocess

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

func trap2(height []int) (ans int) {
	n := len(height)
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for j := n - 2; j >= 0; j-- {
		rightMax[j] = max(rightMax[j+1], height[j])
	}
	for i := 0; i < n; i++ {
		high := min(leftMax[i], rightMax[i])
		if high > height[i] {
			ans += high - height[i]
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
