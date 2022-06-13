package arraymanipulate

import "sort"

// https://leetcode.cn/problems/height-checker/submissions/

func heightChecker(heights []int) (ans int) {
	n := len(heights)
	origin := make([]int, n)
	copy(origin, heights)
	sort.Ints(heights)
	for i := 0; i < n; i++ {
		if origin[i] != heights[i] {
			ans++
		}
	}
	return
}
