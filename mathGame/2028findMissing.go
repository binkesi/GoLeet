package mathgame

// https://leetcode-cn.com/problems/find-missing-observations/

func missingRolls(rolls []int, mean int, n int) (ans []int) {
	l, cur := len(rolls), 0
	for _, v := range rolls {
		cur += v
	}
	left := mean*(l+n) - cur
	if left < n || left > 6*n {
		return
	}
	if left%n == 0 {
		ele := left / n
		for j := 0; j < n; j++ {
			ans = append(ans, ele)
		}
	} else {
		ele := left / n
		for left != 0 {
			if left != n*(ele+1) {
				ans = append(ans, ele)
				left -= ele
			} else {
				ans = append(ans, ele+1)
				left -= ele + 1
			}
			n--
		}
	}
	return
}
