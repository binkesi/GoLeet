package stringparse

// https://leetcode-cn.com/problems/longest-uncommon-subsequence-i/

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	m, n := len(a), len(b)
	if m >= n {
		return m
	} else {
		return n
	}
}
