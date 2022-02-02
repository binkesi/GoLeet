package stringparse

import (
	"strings"
)

// https://leetcode-cn.com/problems/longest-nice-substring/

func longestNiceSubstring(s string) string {
	var ans string
	lenS := len(s)
	for i := 2; i <= lenS; i++ {
		for j := lenS - i; j >= 0; j-- {
			if isNice(s[j : j+i]) {
				ans = s[j : j+i]
			}
		}
	}
	return ans
}

func isNice(s string) bool {
	sMap := make(map[string]struct{})
	for _, v := range s {
		if _, ok := sMap[string(v)]; !ok {
			sMap[string(v)] = struct{}{}
		}
	}
	for i := range sMap {
		if _, ok := sMap[strings.ToLower(i)]; !ok {
			return false
		}
		if _, ok := sMap[strings.ToUpper(i)]; !ok {
			return false
		}
	}
	return true
}
