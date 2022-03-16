package dynamicprocess

// https://leetcode-cn.com/problems/longest-valid-parentheses/

func longestValidParentheses(s string) (ans int) {
	n := len(s)
	if n <= 1 {
		return 0
	}
	var pair func(a, b byte) bool
	pair = func(a, b byte) bool {
		if a == '(' && b == ')' {
			return true
		} else {
			return false
		}
	}
	dp := make([]int, n)
	dp[0] = 0
	for i := 1; i < n; i++ {
		pre := dp[i-1]
		if pre == i {
			dp[i] = 0
		} else if pair(s[i-1-pre], s[i]) {
			if i-2-pre >= 0 {
				dp[i] = dp[i-2-pre] + dp[i-1] + 2
			} else {
				dp[i] = dp[i-1] + 2
			}
		} else {
			dp[i] = 0
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return
}
