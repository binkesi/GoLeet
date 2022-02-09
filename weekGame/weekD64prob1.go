package weekgame

import "sort"

// https://leetcode-cn.com/contest/biweekly-contest-64/problems/kth-distinct-string-in-an-array/

func kthDistinct(arr []string, k int) string {
	type strArr struct {
		str string
		idx int
	}
	arrMap := make(map[string]map[int]int)
	for i, v := range arr {
		if _, ok := arrMap[v]; ok {
			for i := range arrMap[v] {
				arrMap[v][i] += 1
			}
		} else {
			arrMap[v] = make(map[int]int)
			arrMap[v][i] = 1
		}
	}
	arrs := make([]strArr, 0)
	for key, v := range arrMap {
		for i, val := range v {
			if val == 1 {
				arrs = append(arrs, strArr{str: key, idx: i})
			}
		}
	}
	if k > len(arrs) {
		return ""
	}
	sort.Slice(arrs, func(i, j int) bool { return arrs[i].idx < arrs[j].idx })
	return arrs[k-1].str
}
