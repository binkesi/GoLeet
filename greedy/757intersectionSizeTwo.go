package greedy

import "sort"

// https://leetcode.cn/problems/set-intersection-size-at-least-two/

func intersectionSizeTwo(intervals [][]int) (ans int) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return false
		}
	})
	n := len(intervals)
	prepre, pre := intervals[n-1][0], intervals[n-1][0]+1
	ans += 2
	for i := n - 2; i >= 0; i-- {
		if intervals[i][1] < prepre {
			prepre, pre = intervals[i][0], intervals[i][0]+1
			ans += 2
		} else if intervals[i][1] == prepre {
			prepre, pre = intervals[i][0], prepre
			ans += 1
		} else if intervals[i][1] < pre {
			prepre, pre = intervals[i][0], prepre
			ans += 1
		}
	}
	return
}
