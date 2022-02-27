package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-282/problems/minimum-time-to-finish-the-race/

func minimumFinishTime(tires [][]int, changeTime int, numLaps int) int {
	lenT := len(tires)
	sort.Slice(tires, func(i, j int) bool { return tires[i][0] < tires[j][0] })
	use1, use2 := [2]int{}, [2]int{}
	ind1 := 0
	for i := 1; i < lenT; i++ {
		if tires[i][0] != tires[i-1][0] {
			break
		}
		if tires[i][1] < tires[i-1][1] {
			ind1 = i
		}
	}
	copy(use1[:], tires[ind1])
	sort.Slice(tires, func(i, j int) bool { return tires[i][1] < tires[j][1] })
	ind2 := 0
	for i := 1; i < lenT; i++ {
		if tires[i][1] != tires[i-1][0] {
			break
		}
		if tires[i][1] < tires[i-1][1] {
			ind1 = i
		}
	}
}
