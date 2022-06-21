package stringparse

import "strings"

// https://leetcode.cn/problems/defanging-an-ip-address/submissions/

func defangIPaddr(address string) string {
	strs := strings.Split(address, ".")
	res := strings.Join(strs, "[.]")
	return res
}
