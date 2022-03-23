package stringparse

// https://leetcode-cn.com/problems/get-equal-substrings-within-budget/

func equalSubstring(s string, t string, maxCost int) (ans int) {
	n := len(s)
	sDif := make([]int, n)
	for i := 0; i < n; i++ {
		sDif[i] = abs(int(s[i]) - int(t[i]))
	}
	for i := 1; i < n; i++ {
		sDif[i] = sDif[i-1] + sDif[i]
	}
	sDif = append([]int{0}, sDif...)
	for i, j := 0, 1; j < n+1; {
		sum := sDif[j] - sDif[i]
		if sum <= maxCost {
			if j-i > ans {
				ans = j - i
			}
			j++
		} else {
			i++
		}
	}
	return
}
