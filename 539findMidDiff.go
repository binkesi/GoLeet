package goleet

import (
	"math"
	"sort"
	"strconv"
)

// https://leetcode-cn.com/problems/minimum-time-difference/

func findMinDifference(timePoints []string) int {
	lenF := len(timePoints)
	if lenF >= 1440 {
		return 0
	}
	var res int = math.MaxInt32
	var minPoints []int = make([]int, 0)
	const dayMinutes int = 1440
	for i := 0; i < lenF; i++ {
		minPoints = append(minPoints, tranToMin(timePoints[i]))
	}
	sort.Ints(minPoints)
	for k := 1; k < lenF; k++ {
		diff := minPoints[k] - minPoints[k-1]
		if diff < res {
			res = diff
		}
	}
	headTail := dayMinutes - minPoints[lenF-1] + minPoints[0]
	if headTail < res {
		res = headTail
	}
	return res
}

func tranToMin(hourMin string) int {
	hour, _ := strconv.Atoi(hourMin[0:2])
	min, _ := strconv.Atoi(hourMin[3:])
	return hour*60 + min
}
