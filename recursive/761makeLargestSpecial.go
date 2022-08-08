package recursive

import (
	"sort"
	"strings"
)

// https://leetcode.cn/problems/special-binary-string/

func makeLargestSpecial(s string) string {
	var exchange func(str string, line int) string
	exchange = func(str string, line int) string {
		cur, l := 0, len(str)
		if l == 2 {
			return str
		}
		res := []string{}
		start, end := -1, -1
		for i := 0; i < l; i++ {
			if cur == line && end == -1 {
				start = i
			}
			if str[i] == '1' {
				cur++
			} else {
				cur--
			}
			if cur == line && start != -1 {
				end = i
				res = append(res, exchange(str[start:end+1], 1))
				start, end = -1, -1
			}
		}
		sort.Slice(res, func(i, j int) bool { return res[i] > res[j] })
		if line == 0 {
			return strings.Join(res, "")
		} else {
			return str[0:1] + strings.Join(res, "") + str[l-1:l]
		}
	}
	return exchange(s, 0)
}
