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
		num := left % n
		for i := 0; i < n; i++ {
			if i < num {
				ans = append(ans, ele+1)
			} else {
				ans = append(ans, ele)
			}
		}
	}
	return
}
