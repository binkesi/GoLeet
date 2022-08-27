package hashmap

// https://leetcode.cn/problems/make-two-arrays-equal-by-reversing-sub-arrays/

func canBeEqual(target []int, arr []int) bool {
	nMap := make(map[int]int)
	for _, v := range target {
		nMap[v]++
	}
	for _, v := range arr {
		if _, ok := nMap[v]; !ok || nMap[v] <= 0 {
			return false
		}
		nMap[v]--
	}
	return true
}
