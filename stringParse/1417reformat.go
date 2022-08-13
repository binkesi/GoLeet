package stringparse

import "strings"

// https://leetcode.cn/problems/reformat-the-string/

func reformat(s string) string {
	strn, strs := []byte{}, []byte{}
	ans := strings.Builder{}
	for i := range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			strs = append(strs, s[i])
		} else {
			strn = append(strn, s[i])
		}
	}
	if abs(len(strn)-len(strs)) > 1 {
		return ""
	}
	if len(strn)-len(strs) == 1 || len(strn) == len(strs) {
		l := len(strs)
		for i := 0; i < l; i++ {
			ans.WriteByte(strn[i])
			ans.WriteByte(strs[i])
		}
		if len(strn) > len(strs) {
			ans.WriteByte(strn[l])
		}
	} else if len(strs)-len(strn) == 1 {
		l := len(strn)
		for i := 0; i < l; i++ {
			ans.WriteByte(strs[i])
			ans.WriteByte(strn[i])
		}
		ans.WriteByte(strs[l])
	}
	return ans.String()
}
