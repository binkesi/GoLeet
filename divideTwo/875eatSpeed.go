package dividetwo

import "sort"

// https://leetcode.cn/problems/koko-eating-bananas/submissions/

func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	sort.Ints(piles)
	l, r := 1, piles[n-1]
	for l < r {
		mid := l + (r-l)>>1
		cost := 0
		for _, v := range piles {
			if v%mid == 0 {
				cost += v / mid
			} else {
				cost = cost + v/mid + 1
			}
		}
		if cost <= h {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
