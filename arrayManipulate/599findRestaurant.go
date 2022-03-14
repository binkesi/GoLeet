package arraymanipulate

import "math"

// https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/

func findRestaurant(list1 []string, list2 []string) (ans []string) {
	sMap := make(map[string]int)
	for i, s := range list1 {
		sMap[s] = i
	}
	res := math.MaxInt32
	for i, s := range list2 {
		if v, ok := sMap[s]; ok {
			if i+v < res {
				ans = []string{}
				ans = append(ans, s)
				res = i + v
			} else if i+v == res {
				ans = append(ans, s)
			}
		}
	}
	return
}
