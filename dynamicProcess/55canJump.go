package dynamicprocess

// https://leetcode-cn.com/problems/jump-game/

func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if dp[j] == true && nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
	}
	return dp[n-1]
}

func canJump2(nums []int) bool {
	maxInd := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if maxInd >= n-1 {
			return true
		}
		if i > maxInd {
			return false
		} else {
			if nums[i]+i > maxInd {
				maxInd = nums[i] + i
			}
		}
	}
	return false
}
