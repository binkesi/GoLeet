package mathgame

import "math"

// https://leetcode-cn.com/problems/factorial-trailing-zeroes/

func trailingZeroes(n int) (ans int) {
	for i := 1; ; i++ {
		res := n / int(math.Pow(5.0, float64(i)))
		if res > 0 {
			ans += res
		} else {
			break
		}
	}
	return
}
