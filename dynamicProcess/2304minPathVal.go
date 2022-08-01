package dynamicprocess

import "math"

// https://leetcode.cn/problems/minimum-path-cost-in-a-grid/submissions/

func minPathCost(grid [][]int, moveCost [][]int) int {
	r, c := len(grid), len(grid[0])
	dp := make([]int, c)
	copy(dp, grid[0])
	for i := 1; i < r; i++ {
		dpT := make([]int, c)
		copy(dpT, dp)
		for j := 0; j < c; j++ {
			minVal := math.MaxInt32
			for j1 := 0; j1 < c; j1++ {
				if dpT[j1]+moveCost[grid[i-1][j1]][j] < minVal {
					minVal = dpT[j1] + moveCost[grid[i-1][j1]][j]
				}
			}
			dp[j] = minVal + grid[i][j]
		}
	}
	ans := math.MaxInt32
	for i := 0; i < c; i++ {
		if dp[i] < ans {
			ans = dp[i]
		}
	}
	return ans
}
