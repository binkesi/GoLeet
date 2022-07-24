package arraymanipulate

// https://leetcode.cn/problems/distance-between-bus-stops/submissions/

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	disA, disB := 0, 0
	up := max(start, destination)
	for i := min(start, destination); i < up; i++ {
		disA += distance[i]
	}
	for i := range distance {
		disB += distance[i]
	}
	return min(disA, disB-disA)
}
