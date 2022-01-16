package goleet

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-276/problems/maximum-running-time-of-n-computers/

func maxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)
	lenB := len(batteries)
	upper := batteries[lenB-1]*lenB/n + 1
	lower := batteries[lenB-n] - 1
	for lower < upper-1 {
		day := lower + (upper-lower)/2
		sumB := sums(batteries, day)
		if sumB == day*n {
			return int64(day)
		} else if sumB > day*n {
			lower = day
		} else {
			upper = day
		}
	}
	return int64(lower)
}

func sums(nums []int, n int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] <= n {
			res += nums[i]
		} else {
			res += n
		}
	}
	return res
}
