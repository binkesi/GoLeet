package weekgame

import "unicode"

// https://leetcode.cn/contest/weekly-contest-298/problems/greatest-english-letter-in-upper-and-lower-case/

func greatestLetter(s string) string {
	bMap := make(map[byte]int)
	var res string
	for _, b := range s {
		if unicode.IsLower(b) {
			if _, ok := bMap[byte(unicode.ToUpper(b))]; !ok {
				bMap[byte(unicode.ToUpper(b))] = 0
			} else {
				if bMap[byte(unicode.ToUpper(b))] == 1 {
					bMap[byte(unicode.ToUpper(b))] = 2
				}
			}
		} else if unicode.IsUpper(b) {
			if _, ok := bMap[byte(b)]; !ok {
				bMap[byte(b)] = 1
			} else {
				if bMap[byte(b)] == 0 {
					bMap[byte(b)] = 2
				}
			}
		}
	}
	for k, v := range bMap {
		if v == 2 {
			if res == "" {
				res = string(k)
			} else if k > byte(res[0]) {
				res = string(k)
			}
		}
	}
	return string(res)
}
