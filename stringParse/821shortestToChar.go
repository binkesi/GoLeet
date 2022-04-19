package stringparse

// https://leetcode-cn.com/problems/shortest-distance-to-a-character/

func shortestToChar(s string, c byte) (ans []int) {
	cind := make([]int, 0)
	for i, b := range s {
		if byte(b) == c {
			cind = append(cind, i)
		}
	}
	cur, cnt := 0, len(cind)
	for i, b := range s {
		if byte(b) == c {
			cur++
			ans = append(ans, 0)
			continue
		}
		if cur == 0 {
			ans = append(ans, abs(i-cind[0]))
		} else if cur == cnt {
			ans = append(ans, abs(i-cind[cnt-1]))
		} else {
			ans = append(ans, min(abs(i-cind[cur-1]), abs(i-cind[cur])))
		}
	}
	return
}
