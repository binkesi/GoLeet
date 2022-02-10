package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-196/problems/last-moment-before-all-ants-fall-out-of-a-plank/

func getLastMoment(n int, left []int, right []int) int {
	sort.Ints(left)
	sort.Ints(right)
	l, r := len(left), len(right)
	if l == 0 {
		return n - right[0]
	}
	if r == 0 {
		return left[l-1]
	}
	if left[l-1] > (n - right[0]) {
		return left[l-1]
	} else {
		return n - right[0]
	}
}
