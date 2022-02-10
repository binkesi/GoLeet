package weekgame

import "sort"

// https://leetcode-cn.com/contest/biweekly-contest-64/problems/two-best-non-overlapping-events/

func maxTwoEvents(events [][]int) int {
	lenE := len(events)
	dp := make([][2]int, lenE)
	sort.Slice(events, func(i, j int) bool { return events[i][1] < events[j][1] })
	dp[0][0], dp[0][1] = events[0][2], events[0][2]
	for i := 1; i < lenE; i++ {
		idx := -1
		x, val := events[i][0], events[i][2]
		for l, h := 0, i-1; l <= h; {
			mid := l + (h-l)/2
			if events[mid][1] < x && events[mid+1][1] >= x {
				idx = mid
				break
			} else if events[mid][1] < x {
				l = mid + 1
			} else if events[mid][1] >= x {
				h = mid - 1
			}
		}
		if idx >= 0 {
			dp[i][0] = maxN(dp[i-1][0], val)
			dp[i][1] = maxN(dp[i-1][1], dp[idx][0]+val)
		} else {
			dp[i][0] = maxN(dp[i-1][0], val)
			dp[i][1] = maxN(dp[i-1][1], val)
		}
	}
	return dp[lenE-1][1]
}

func maxN(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
