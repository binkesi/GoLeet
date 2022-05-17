package stringparse

// https://leetcode.cn/problems/verifying-an-alien-dictionary/submissions/

func isAlienSorted(words []string, order string) bool {
	sMap := make(map[byte]int)
	for i, v := range order {
		sMap[byte(v)] = i
	}
	var low func(a, b string) bool
	low = func(a, b string) bool {
		m, n := len(a), len(b)
		for i := 0; i < m && i < n; i++ {
			if sMap[a[i]] > sMap[b[i]] {
				return false
			}
			if sMap[a[i]] < sMap[b[i]] {
				return true
			}
		}
		return m <= n
	}
	n := len(words)
	for i := 0; i < n-1; i++ {
		if !low(words[i], words[i+1]) {
			return false
		}
	}
	return true
}
