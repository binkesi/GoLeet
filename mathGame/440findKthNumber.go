package mathgame

import (
	"math"
	"strconv"
)

// https://leetcode-cn.com/problems/k-th-smallest-in-lexicographical-order/

func findKthNumber(n int, k int) int {
	ans := 1
	for k > 1 {
		cnt := getCnt(ans, n)
		if cnt < k {
			k -= cnt
			ans++
		} else {
			k--
			ans *= 10
		}
	}
	return ans
}

func getCnt(x, limit int) int {
	a, b := strconv.Itoa(x), strconv.Itoa(limit)
	n, m := len(a), len(b)
	k, ans := m-n, 0
	u, _ := strconv.Atoi(b[:n])
	for i := 0; i < k; i++ {
		ans += int(math.Pow10(i))
	}
	if u > x {
		ans += int(math.Pow10(k))
	} else if u == x {
		ans += (limit - x*int(math.Pow10(k)) + 1)
	}
	return ans
}
