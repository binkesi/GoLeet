package arraymanipulate

import (
	"sort"
)

// https://leetcode-cn.com/problems/the-number-of-weak-characters-in-the-game/

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		p, q := properties[i], properties[j]
		return p[0] > q[0] || p[0] == q[0] && p[1] < q[1]
	})
	maxdf, ans := 0, 0
	for _, v := range properties {
		if v[1] < maxdf {
			ans += 1
		} else {
			maxdf = v[1]
		}
	}
	return ans
}
