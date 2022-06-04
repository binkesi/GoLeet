package stringparse

import (
	"strings"
)

// https://leetcode.cn/problems/unique-email-addresses/

func numUniqueEmails(emails []string) int {
	eMap := make(map[string]struct{})
	for _, v := range emails {
		addr := parse(v)
		if _, ok := eMap[addr]; !ok {
			eMap[addr] = struct{}{}
		}
	}
	return len(eMap)
}

func parse(s string) string {
	start := 0
	for i, v := range s {
		if v == '@' {
			if start != 0 {
				s = s[0:start] + s[i:]
				s = strings.Join(strings.Split(s[0:start], "."), "") + s[start:]
			} else {
				s = strings.Join(strings.Split(s[0:i], "."), "") + s[i:]
			}
			break
		}
		if v == '+' && start == 0 {
			start = i
		}
	}
	return s
}
