package mathgame

import "strconv"

// https://leetcode-cn.com/problems/simplified-fractions/

func simplifiedFractions(n int) (ans []string) {
	res := [][2]int{}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if gcd(i, j) == 1 {
				res = append(res, [2]int{j, i})
			}
		}
	}
	for _, v := range res {
		str := strconv.Itoa(v[0]) + "/" + strconv.Itoa(v[1])
		ans = append(ans, str)
	}
	return ans
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return b
}
