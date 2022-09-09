package stringparse

import "strings"

// https://leetcode.cn/problems/crawler-log-folder/submissions/

func minOperations(logs []string) (ans int) {
	for _, v := range logs {
		if strings.Count(v, ".") == 0 {
			ans++
		}
		if strings.Count(v, ".") == 2 && ans > 0 {
			ans--
		}
	}
	return
}
