package greedy

import "sort"

// https://leetcode.cn/problems/max-chunks-to-make-sorted-ii/submissions/

func maxChunksToSorted(arr []int) (ans int) {
	l := len(arr)
	sortArr := make([]int, l)
	copy(sortArr, arr)
	sort.Ints(sortArr)
	nMap := make(map[int]int)
	sum := 0
	for i := 0; i < l; i++ {
		if nMap[sortArr[i]] < 0 {
			sum--
		}
		nMap[sortArr[i]]++
		if nMap[arr[i]] > 0 {
			sum--
		}
		nMap[arr[i]]--
		sum++
		if sum == 0 {
			ans++
		}
	}
	return
}
