package mathgame

// https://leetcode.cn/problems/beautiful-arrangement-ii/

func constructArray(n int, k int) (ans []int) {
	for i := 1; i <= n; i++ {
		ans = append(ans, i)
	}
	if n == 1 {
		return
	}
	l, r := n-2, n-1
	for j := k; j > 1; j-- {
		tmpa, tmpb := l, r
		for tmpb < n {
			ans[tmpa], ans[tmpb] = ans[tmpb], ans[tmpa]
			tmpa += 2
			tmpb += 2
		}
		l--
		r--
	}
	return
}

func constructArray2(n int, k int) (ans []int) {
	for i := 1; i <= n-k; i++ {
		ans = append(ans, i)
	}
	l, r := n-k+1, n
	for l <= r {
		if l == r {
			ans = append(ans, l)
		} else {
			ans = append(ans, r)
			ans = append(ans, l)
		}
		l++
		r--
	}
	return
}
