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
		if arr[i]-arr[i-1] < minDiff {
			ans = [][]int{}
			ans = append(ans, []int{arr[i-1], arr[i]})
			minDiff = arr[i] - arr[i-1]
		} else if arr[i]-arr[i-1] == minDiff {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
	return
}
