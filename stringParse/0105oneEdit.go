package stringparse

// https://leetcode.cn/problems/one-away-lcci/submissions/

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m-n >= 2 || n-m >= 2 {
		return false
	}
	if m == n {
		cnt := 0
		for i := 0; i < m; i++ {
			if first[i] != second[i] {
				cnt++
			}
		}
		return cnt <= 1
	}
	if m < n {
		first, second = second, first
		m, n = n, m
	}
	cnt := 0
	for i, j := 0, 0; j < n; i++ {
		if cnt == 0 && first[i] != second[j] {
			cnt++
			continue
		} else if first[i] != second[j] {
			return false
		}
		j++
	}
	return cnt <= 1
}
