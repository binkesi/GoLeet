package stringparse

// https://leetcode-cn.com/problems/rotate-string/

func rotateString(s string, goal string) bool {
	n := len(s)
	if n != len(goal) {
		return false
	}
	for i := 0; i < n; i++ {
		if s[i] == goal[0] {
			if s[0:i] == goal[n-i:n] && s[i:] == goal[0:n-i] {
				return true
			}
		}
	}
	return false
}
