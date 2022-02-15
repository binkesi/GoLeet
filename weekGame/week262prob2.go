package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-262/problems/minimum-operations-to-make-a-uni-value-grid/

func minOperations(grid [][]int, x int) (ans int) {
	nums := []int{}
	r, c := len(grid), len(grid[0])
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			nums = append(nums, grid[i][j])
		}
	}
	sort.Ints(nums)
	lenN := len(nums)
	mid := nums[lenN/2]
	ans += abs(nums[0]-mid) / x
	for i := 1; i < lenN; i++ {
		if (nums[i]-nums[i-1])%x != 0 {
			return -1
		}
		ans += abs(nums[i]-mid) / x
	}
	return
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
