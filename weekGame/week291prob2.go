package weekgame

import "math"

// https://leetcode.cn/contest/weekly-contest-291/problems/minimum-consecutive-cards-to-pick-up/

func minimumCardPickup(cards []int) int {
	res := math.MaxInt32
	cMap := make(map[int][]int)
	for i := range cards {
		cMap[cards[i]] = append(cMap[cards[i]], i)
	}
	for _, v := range cMap {
		n := len(v)
		if n > 1 {
			for p1, p2 := 0, 1; p2 < n; {
				if v[p2]-v[p1] < res {
					res = v[p2] - v[p1]
				}
				p1++
				p2++
			}
		}
	}
	if res != math.MaxInt32 {
		return res + 1
	} else {
		return -1
	}
}
