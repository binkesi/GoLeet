package weekgame

import "sort"

// https://leetcode-cn.com/contest/biweekly-contest-71/problems/minimum-sum-of-four-digit-number-after-splitting-digits/

func minimumSum(num int) (ans int) {
	res := []int{}
	for num > 0 {
		res = append(res, num%10)
		num /= 10
	}
	sort.Ints(res)
	ans = (res[0]+res[1])*10 + res[2] + res[3]
	return
}
