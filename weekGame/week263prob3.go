package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-263/problems/count-number-of-maximum-bitwise-or-subsets/

func countMaxOrSubsets(nums []int) int {
	var res int
	var target int
	n := len(nums)
	for _, num := range nums {
		target |= num
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	var dfs func(idx, curr int)
	dfs = func(idx, curr int) {
		if curr == target {
			res += 1 << (n - idx)
			return
		}
		if idx == n {
			return
		}
		dfs(idx+1, curr|nums[idx])
		dfs(idx+1, curr)
	}
	dfs(0, 0)
	return res
}
