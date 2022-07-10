package dynamicprocess

import (
	"math"
)

// https://leetcode.cn/problems/cherry-pickup/submissions/

func cherryPickup(grid [][]int) int {
	r := len(grid)
	dp := make([][][]int, 2*r+1)
	for i := range dp {
		dp[i] = make([][]int, r+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, r+1)
		}
	}
	for k := 0; k <= 2*r; k++ {
		for i1 := 0; i1 <= r; i1++ {
			for i2 := 0; i2 <= r; i2++ {
				dp[k][i1][i2] = math.MinInt32
			}
		}
	}
	dp[2][1][1] = grid[0][0]
	for k := 3; k <= 2*r; k++ {
		for x1 := 1; x1 <= r; x1++ {
			for x2 := 1; x2 <= r; x2++ {
				y1, y2 := k-x1, k-x2
				if y1 <= 0 || y1 > r || y2 <= 0 || y2 > r {
					continue
				}
				if grid[x1-1][y1-1] == -1 || grid[x2-1][y2-1] == -1 {
					continue
				}
				tmp := max(max(dp[k-1][x1-1][x2], dp[k-1][x1-1][x2-1]), max(dp[k-1][x1][x2], dp[k-1][x1][x2-1]))
				if x1 == x2 {
					dp[k][x1][x2] = tmp + grid[x1-1][y1-1]
				} else {
					dp[k][x1][x2] = tmp + grid[x1-1][y1-1] + grid[x2-1][y2-1]
				}
			}
		}
	}
	if dp[2*r][r][r] <= 0 {
		return 0
	}
	return dp[2*r][r][r]
}
