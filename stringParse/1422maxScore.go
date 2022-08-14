package stringparse

// https://leetcode.cn/problems/maximum-score-after-splitting-a-string/

func maxScore(s string) int {
	sc, cntOne, cntZero := 0, 0, 0
	for i := range s {
		if s[i] == '0' {
			cntZero++
		} else {
			cntOne++
		}
	}
	cur := cntOne
	for i := range s {
		if i == len(s)-1 {
			break
		}
		if s[i] == '0' {
			cur++
		} else {
			cur--
		}
		if i == 0 {
			sc = cur
		}
		if cur > sc {
			sc = cur
		}
	}
	return sc
}
