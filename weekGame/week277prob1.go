package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-277/problems/count-elements-with-strictly-smaller-and-greater-elements/

func countElements(nums []int) int {
	lenN := len(nums)
	if lenN <= 2 {
		return 0
	}
	sort.Ints(nums)
	low, high := nums[0], nums[lenN-1]
	var res int = 0
	for i := 1; i < lenN-1; i++ {
		if nums[i] > low && nums[i] < high {
			res += 1
		}
	}
	return res
}
