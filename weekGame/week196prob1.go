package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-196/problems/can-make-arithmetic-progression-from-sequence/

func canMakeArithmeticProgression(arr []int) bool {
	ans := true
	lenA := len(arr)
	if lenA <= 1 {
		return ans
	}
	sort.Ints(arr)
	minus := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != minus {
			ans = false
			break
		}
	}
	return ans
}
