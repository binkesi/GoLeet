package arraymanipulate

import "sort"

// https://leetcode.cn/problems/wiggle-sort-ii/

func wiggleSort(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	sort.Ints(nums)
	tmp := make([]int, n)
	for i, j, k := 0, (n-1)/2, n-1; k > (n-1)/2; {
		tmp[i] = nums[j]
		i++
		tmp[i] = nums[k]
		i++
		j--
		k--
		if n%2 == 1 {
			tmp[i] = nums[0]
		}
	}
	copy(nums, tmp)
}
