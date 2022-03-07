package mathgame

import "strconv"

// https://leetcode-cn.com/problems/base-7/

func convertToBase7(num int) (ans string) {
	if num == 0 {
		ans = "0"
		return
	}
	tmp := ab(num)
	for tmp != 0 {
		ans = strconv.Itoa(tmp%7) + ans
		tmp = tmp / 7
	}
	if num < 0 {
		ans = "-" + ans
	}
	return
}

func ab(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
