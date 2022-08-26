package selectsort

import "sort"

// https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array/

func maxProduct(nums []int) int {
	sort.Ints(nums)
	l := len(nums)
	return max((nums[0]-1)*(nums[1]-1), (nums[l-1]-1)*(nums[l-2]-1))
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
