package arraymanipulate

import "sort"

// https://leetcode-cn.com/problems/merge-intervals/

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	ans = append(ans, intervals[0])
	n := len(intervals)
	for i := 1; i < n; i++ {
		la := len(ans)
		p1, p2 := intervals[i][0], intervals[i][1]
		lp2 := ans[la-1][1]
		if p1 > lp2 {
			ans = append(ans, intervals[i])
		} else if p2 > lp2 {
			ans[la-1][1] = p2
		}
	}
	return ans
}
