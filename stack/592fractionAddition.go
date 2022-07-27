package stack

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/fraction-addition-and-subtraction/

func fractionAddition(expression string) string {
	times := 2520
	divid, divided := make([]int, 0), make([]int, 0)
	symbol, l := 1, len(expression)
	for i, b := range expression {
		if b == '-' {
			symbol = -1
		}
		if b == '+' {
			symbol = 1
		}
		if b == '/' {
			if expression[i-1] == '0' {
				divid = append(divid, 10*symbol)
			} else {
				divid = append(divid, int(expression[i-1]-'0')*symbol)
			}
			if expression[i+1] == '1' && i < l-2 && expression[i+2] == '0' {
				divided = append(divided, 10)
			} else {
				divided = append(divided, int(expression[i+1]-'0'))
			}
		}
	}
	sum := 0
	for i, v := range divided {
		mul := times / v
		sum += divid[i] * mul
	}
	did := gcd(abs(sum), times)
	s := strings.Builder{}
	s.WriteString(strconv.Itoa(sum / did))
	s.WriteByte('/')
	s.WriteString(strconv.Itoa(times / did))
	return s.String()
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	a, b = b, a%b
	return gcd(a, b)
}
