package greedy

// https://leetcode-cn.com/problems/longest-happy-string/

func longestDiverseString(a int, b int, c int) (ans string) {
	curArr := [2]string{}
	for a > 0 || b > 0 || c > 0 {
		if a == max(a, b, c) {
			if curArr[0] != "a" || curArr[1] != "a" {
				ans += "a"
				curArr[0] = curArr[1]
				curArr[1] = "a"
				a -= 1
			} else if b >= c && b > 0 {
				ans += "b"
				curArr[0] = curArr[1]
				curArr[1] = "b"
				b -= 1
			} else if c > 0 {
				ans += "c"
				curArr[0] = curArr[1]
				curArr[1] = "c"
				c -= 1
			} else {
				return
			}
		} else if b == max(a, b, c) {
			if curArr[0] != "b" || curArr[1] != "b" {
				ans += "b"
				curArr[0] = curArr[1]
				curArr[1] = "b"
				b -= 1
			} else if a >= c && a > 0 {
				ans += "a"
				curArr[0] = curArr[1]
				curArr[1] = "a"
				a -= 1
			} else if c > 0 {
				ans += "c"
				curArr[0] = curArr[1]
				curArr[1] = "c"
				c -= 1
			} else {
				return
			}
		} else if c == max(a, b, c) {
			if curArr[0] != "c" || curArr[1] != "c" {
				ans += "c"
				curArr[0] = curArr[1]
				curArr[1] = "c"
				c -= 1
			} else if a >= b && a > 0 {
				ans += "a"
				curArr[0] = curArr[1]
				curArr[1] = "a"
				a -= 1
			} else if b > 0 {
				ans += "b"
				curArr[0] = curArr[1]
				curArr[1] = "b"
				b -= 1
			} else {
				return
			}
		}
	}
	return
}

func max(a, b, c int) int {
	if a >= b && a >= c {
		return a
	} else if b >= c {
		return b
	} else {
		return c
	}
}
