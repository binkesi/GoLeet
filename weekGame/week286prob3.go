package weekgame

import (
	"fmt"
	"math"
	"strconv"
)

// https://leetcode-cn.com/contest/weekly-contest-286/problems/find-palindrome-with-fixed-length/

func kthPalindrome(queries []int, intLength int) (ans []int64) {
	n := len(queries)
	var total int
	if intLength == 1 {
		total = 9
		for i := 0; i < n; i++ {
			if queries[i] > total {
				ans = append(ans, -1)
			} else {
				ans = append(ans, int64(queries[i]))
			}
		}
		return
	}
	if intLength == 2 {
		total = 9
		for i := 0; i < n; i++ {
			if queries[i] > total {
				ans = append(ans, -1)
			} else {
				ans = append(ans, int64(queries[i]*11))
			}
		}
		return
	}
	total = 9 * int(math.Pow10((intLength-1)/2))
	fmt.Println(total)
	for i := 0; i < n; i++ {
		tmp := queries[i]
		if tmp > total {
			ans = append(ans, -1)
		} else {
			strs := ""
			for j := 0; j <= (intLength-1)/2; j++ {
				if j == 0 {
					res := (tmp-1)/int(math.Pow10((intLength-1)/2)) + 1
					strs += strconv.Itoa(res)
					tmp -= (res-1)*int(math.Pow10((intLength-1)/2)) + 1
				} else {
					res := tmp / int(math.Pow10((intLength-1)/2-j))
					if len(strs) != 0 {
						strs += strconv.Itoa(res)
						tmp -= res * int(math.Pow10((intLength-1)/2-j))
					} else {
						if res > 0 {
							strs += strconv.Itoa(res)
							tmp %= int(math.Pow10((intLength-1)/2 - j))
						}
					}
				}
			}
			l := len(strs)
			for k := l - 1; k >= 0; k-- {
				if k == l-1 {
					if intLength%2 == 1 {
						continue
					}
				}
				strs += string(strs[k])
			}
			num, _ := strconv.Atoi(strs)
			ans = append(ans, int64(num))
		}
	}
	return
}
