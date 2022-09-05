package greedy

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/maximum-length-of-pair-chain/

func findLongestChain(pairs [][]int) (ans int) {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	curMin := math.MinInt32
	for _, v := range pairs {
		if v[0] > curMin {
			ans++
			curMin = v[1]
		}
	}
	return
}
