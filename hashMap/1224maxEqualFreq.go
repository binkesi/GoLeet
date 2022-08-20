package hashmap

import "sort"

// https://leetcode.cn/problems/maximum-equal-frequency/

func maxEqualFreq(nums []int) int {
	nMap := make(map[int]int)
	for _, v := range nums {
		nMap[v]++
	}
	nArr := make([][2]int, 0)
	for k, v := range nMap {
		nArr = append(nArr, [2]int{k, v})
	}
	sort.Slice(nArr, func(i, j int) bool { return nArr[i][1] < nArr[j][1] })
	for i, v := range nArr {
		nMap[v[0]] = i
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if isEqual(nArr) {
			return i + 1
		}
		num := nums[i]
		ind := nMap[num]
		nArr[ind][1]--
		for ind > 0 && nArr[ind][1] < nArr[ind-1][1] {
			nMap[num]--
			nMap[nArr[ind-1][0]]++
			nArr[ind-1], nArr[ind] = nArr[ind], nArr[ind-1]
			ind--
		}
		if nArr[0][1] == 0 {
			nArr = nArr[1:]
			for _, v := range nArr {
				nMap[v[0]]--
			}
		}
	}
	return 0
}

func isEqual(arr [][2]int) bool {
	l := len(arr)
	if l == 1 {
		return true
	}
	if l == 2 && (arr[0][1] == 1 || arr[1][1] == 1 || arr[0][1]+1 == arr[1][1]) {
		return true
	}
	if (arr[0][1] == 1 && arr[1][1] == arr[l-1][1]) || (arr[0][1] == arr[l-2][1] && arr[l-1][1] == arr[0][1]+1) {
		return true
	}
	return false
}
