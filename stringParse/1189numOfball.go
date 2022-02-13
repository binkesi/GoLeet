package stringparse

import "math"

// https://leetcode-cn.com/problems/maximum-number-of-balloons/

func maxNumberOfBalloons(text string) (ans int) {
	strMap := make(map[rune]int)
	strMap['b'], strMap['a'], strMap['l'], strMap['o'], strMap['n'] = 0, 0, 0, 0, 0
	for _, v := range text {
		if v == 'b' || v == 'a' || v == 'l' || v == 'o' || v == 'n' {
			strMap[v] += 1
		}
	}
	strMap['l'] /= 2
	strMap['o'] /= 2
	ans = math.MaxInt32
	for _, v := range strMap {
		if ans > v {
			ans = v
		}
	}
	return
}
