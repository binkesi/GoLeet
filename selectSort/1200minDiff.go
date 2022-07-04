package selectsort

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/minimum-absolute-difference/submissions/

func minimumAbsDifference(arr []int) (ans [][]int) {
	sort.Ints(arr)
	minDiff := math.MaxInt32
	n := len(arr)
	for i := 1; i < n; i++ {
		if abs(arr[i]-arr[i-1]) < minDiff {
			ans = [][]int{}
			if arr[i] > arr[i-1] {
				ans = append(ans, []int{arr[i-1], arr[i]})
			} else {
				ans = append(ans, []int{arr[i], arr[i-1]})
			}
			minDiff = abs(arr[i] - arr[i-1])
		} else if abs(arr[i]-arr[i-1]) == minDiff {
			if arr[i] > arr[i-1] {
				ans = append(ans, []int{arr[i-1], arr[i]})
			} else {
				ans = append(ans, []int{arr[i], arr[i-1]})
			}
		}
	}
	return
}
