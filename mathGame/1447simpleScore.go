package mathgame

import "strconv"

// https://leetcode-cn.com/problems/simplified-fractions/

func simplifiedFractions(n int) (ans []string) {
	res := [][2]int{}
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			if isValidFrac(j, i) {
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

func isValidFrac(a, b int) bool {
	for i := 2; i <= a && i <= b/2; i++ {
		if a%i == 0 && b%i == 0 {
			return false
		}
	}
	return true
}
