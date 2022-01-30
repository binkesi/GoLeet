package weekgame

// https://leetcode-cn.com/contest/weekly-contest-258/problems/number-of-pairs-of-interchangeable-rectangles/

func interchangeableRectangles(rectangles [][]int) int64 {
	ans := 0
	tmp := []float64{}
	for _, v := range rectangles {
		tmp = append(tmp, float64(v[0])/float64(v[1]))
	}
	resMap := make(map[float64]int)
	for _, v := range tmp {
		if _, ok := resMap[v]; !ok {
			resMap[v] = 1
		} else {
			resMap[v] += 1
		}
	}
	for i := range resMap {
		ans += resMap[i] * (resMap[i] - 1) / 2
	}
	return int64(ans)
}
