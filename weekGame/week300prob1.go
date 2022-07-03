package weekgame

import "strings"

// https://leetcode.cn/contest/weekly-contest-300/problems/decode-the-message/

func decodeMessage(key string, message string) string {
	mesMap := make(map[byte]byte)
	tmp := 0
	for _, v := range key {
		if byte(v) == ' ' {
			continue
		}
		if _, ok := mesMap[byte(v)]; !ok {
			mesMap[byte(v)] = byte(int('a') + tmp)
			tmp++
		} else {
			continue
		}
	}
	ans := strings.Builder{}
	for _, v := range message {
		if byte(v) == ' ' {
			ans.WriteByte(byte(v))
		} else {
			ans.WriteByte(mesMap[(byte(v))])
		}
	}
	return ans.String()
}
