package stringparse

import "strconv"

// https://leetcode-cn.com/problems/complex-number-multiplication/

func complexNumberMultiply(num1 string, num2 string) (ans string) {
	n1, n2 := [2]int{}, [2]int{}
	l1, l2 := len(num1), len(num2)
	for i := 0; i < l1; i++ {
		if num1[i] == '+' {
			n1[0], _ = strconv.Atoi(num1[0:i])
			n1[1], _ = strconv.Atoi(num1[i+1 : l1-1])
			break
		}
	}
	for j := 0; j < l2; j++ {
		if num2[j] == '+' {
			n2[0], _ = strconv.Atoi(num2[0:j])
			n2[1], _ = strconv.Atoi(num2[j+1 : l2-1])
			break
		}
	}
	x := n1[0]*n2[0] - n1[1]*n2[1]
	y := n1[0]*n2[1] + n2[0]*n1[1]
	ans = strconv.Itoa(x) + "+" + strconv.Itoa(y) + "i"
	return
}
