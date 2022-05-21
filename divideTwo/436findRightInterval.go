package dividetwo

import "sort"

// https://leetcode.cn/problems/find-right-interval/submissions/

func findRightInterval(intervals [][]int) (ans []int) {
	iArr := [][2]int{}
	var find func(int, [][2]int) [2]int
	find = func(a int, arr [][2]int) [2]int {
		n := len(arr)
		l, r := 0, n-1
		if arr[r][1] < a {
			return [2]int{-1, -1}
		}
		for r > l {
			mid := l + (r-l)>>1
			if arr[mid][1] < a {
				l = mid + 1
			} else {
				r = mid
			}
		}
		return arr[l]
	}
	for i, v := range intervals {
		iArr = append(iArr, [2]int{i, v[0]})
	}
	sort.Slice(iArr, func(i, j int) bool { return iArr[i][1] < iArr[j][1] })
	for _, v := range intervals {
		res := find(v[1], iArr)
		ans = append(ans, res[0])
	}
	return
}
