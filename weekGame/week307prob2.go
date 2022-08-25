package weekgame

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/contest/weekly-contest-307/problems/largest-palindromic-number/

func largestPalindromic(num string) string {
	numArr := make([]int, 10)
	for _, b := range num {
		numArr[int(b-'0')]++
	}
	res := strings.Builder{}
	maxOdd := -1
	for i := 9; i >= 0; i-- {
		if numArr[i]%2 == 1 && i > maxOdd {
			maxOdd = i
		}
		if i != 0 {
			res.WriteString(strings.Repeat(strconv.Itoa(i), numArr[i]/2))
		} else if len(res.String()) >= 1 {
			res.WriteString(strings.Repeat(strconv.Itoa(i), numArr[i]/2))
		}
	}
	tmp := make([]byte, len(res.String()))
	copy(tmp, []byte(res.String()))
	for i, j := 0, len(tmp)-1; i < j; {
		tmp[i], tmp[j] = tmp[j], tmp[i]
		i++
		j--
	}
	if maxOdd != -1 {
		res.WriteString(strings.Repeat(strconv.Itoa(maxOdd), 1))
	}
	if maxOdd == -1 && len(res.String()) == 0 && numArr[0] > 0 {
		return "0"
	}
	return res.String() + string(tmp)
}
