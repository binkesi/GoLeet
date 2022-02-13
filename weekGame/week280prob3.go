package weekgame

import (
	"math"
	"sort"
)

// https://leetcode-cn.com/contest/weekly-contest-280/problems/removing-minimum-number-of-magic-beans/

func minimumRemoval(beans []int) (ans int64) {
	lenB := len(beans)
	sort.Ints(beans)
	preSum, preMinSum := make([]int, lenB), make([]int, lenB)
	catch := make([]int, lenB)
	preSum[0], preMinSum[0] = beans[0], 0
	for i := 1; i < lenB; i++ {
		preSum[i] = beans[i] + preSum[i-1]
	}
	for i := 1; i < lenB; i++ {
		preMinSum[0] += (beans[i] - beans[0])
	}
	for i := 1; i < lenB; i++ {
		preMinSum[i] = preMinSum[i-1] - (beans[i]-beans[i-1])*(lenB-i)
	}
	for i := 0; i < lenB; i++ {
		if i == 0 {
			catch[i] = preMinSum[i]
		} else {
			catch[i] = preSum[i-1] + preMinSum[i]
		}
	}
	ans = math.MaxInt64
	for _, v := range catch {
		if int64(v) < ans {
			ans = int64(v)
		}
	}
	return
}
