package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-281/problems/construct-string-with-repeat-limit/

func repeatLimitedString(s string, repeatLimit int) string {
	lenS := len(s)
	bMap := make(map[byte]int)
	blis := make([]int, 0)
	res := make([]byte, 0)
	for i := 0; i < lenS; i++ {
		bMap[s[i]] += 1
	}
	for k, _ := range bMap {
		blis = append(blis, int(k))
	}
	sort.Ints(blis)
out:
	for j := len(blis) - 1; j >= 0; j-- {
		curb := byte(blis[j])
		for bMap[curb] > repeatLimit {
			for k := 0; k < repeatLimit; k++ {
				res = append(res, curb)
			}
			if j == 0 {
				break out
			}
			bMap[curb] -= repeatLimit
			for l := j - 1; l >= 0; l-- {
				insertb := byte(blis[l])
				if bMap[insertb] != 0 {
					res = append(res, insertb)
					bMap[insertb] -= 1
					break
				}
				if l == 0 {
					break out
				}
			}
		}
		for i := 0; i < bMap[curb]; i++ {
			res = append(res, curb)
		}
	}
	return string(res)
}
