package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-279/problems/smallest-value-of-the-rearranged-number/

func smallestNumber(num int64) (ans int64) {
	pos := true
	if num < 0 {
		pos = false
		num *= -1
	}
	res := make([]int, 0)
	for num > 0 {
		res = append(res, int(num%10))
		num /= 10
	}
	sort.Ints(res)
	if !pos {
		for i := len(res) - 1; i >= 0; i-- {
			ans *= 10
			ans += int64(res[i])
		}
		return -ans
	} else {
		for i := 0; i < len(res); i++ {
			if res[i] != 0 {
				res[0], res[i] = res[i], res[0]
				break
			}
		}
		for j := 0; j < len(res); j++ {
			ans *= 10
			ans += int64(res[j])
		}
		return ans
	}
}
