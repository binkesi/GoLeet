package twopointer

import "sort"

// https://leetcode-cn.com/problems/count-number-of-pairs-with-absolute-difference-k/

func countKDifference(nums []int, k int) (ans int) {
	sort.Ints(nums)
	lenN := len(nums)
	l, h := 0, 1
	for h < lenN {
		if nums[h]-nums[l] < k {
			h += 1
		} else if nums[h]-nums[l] > k {
			l += 1
		} else {
			cntl, cnth := 1, 1
			for nums[l+1] == nums[l] {
				cntl += 1
				l += 1
			}
			for nums[h+1] == nums[h] {
				cnth += 1
				h += 1
			}
			ans += cntl * cnth
			l += 1
			h += 1
		}
	}
	return
}
