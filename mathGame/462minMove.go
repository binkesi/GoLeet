package mathgame

import "sort"

func minMoves2(nums []int) (ans int) {
	n := len(nums)
	sort.Ints(nums)
	mid := nums[n/2]
	for _, v := range nums {
		ans += abs1(v - mid)
	}
	return
}

func abs1(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
