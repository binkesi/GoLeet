package twopointer

import "sort"

// https://leetcode.cn/problems/k-diff-pairs-in-an-array/submissions/

func findPairs(nums []int, k int) (ans int) {
	n := len(nums)
	sort.Ints(nums)
	l, r := 0, 1
	for r < n {
		if r == l {
			if r < n-1 {
				r++
			} else {
				break
			}
		}
		if nums[r]-nums[l] == k {
			ans++
			for l+1 < n && nums[l+1] == nums[l] {
				l++
			}
			for r+1 < n && nums[r+1] == nums[r] {
				r++
			}
			l++
			r++
		} else if nums[r]-nums[l] < k {
			for r+1 < n && nums[r+1] == nums[r] {
				r++
			}
			r++
		} else if nums[r]-nums[l] > k {
			for l+1 < n && nums[l+1] == nums[l] {
				l++
			}
			l++
		}
	}
	return
}
