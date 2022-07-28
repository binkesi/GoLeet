package selectsort

import "sort"

// https://leetcode.cn/problems/rank-transform-of-an-array/

func arrayRankTransform(arr []int) []int {
	tmp := make([]int, len(arr))
	copy(tmp, arr)
	sort.Ints(tmp)
	numMap := make(map[int]int)
	ind := 1
	for _, v := range tmp {
		if _, ok := numMap[v]; !ok {
			numMap[v] = ind
			ind++
		}
	}
	for i := range arr {
		arr[i] = numMap[arr[i]]
	}
	return arr
}
