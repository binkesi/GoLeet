package arraymanipulate

// https://leetcode.cn/problems/minimum-cost-to-move-chips-to-the-same-position/submissions/

func minCostToMoveChips(position []int) int {
	p1, p2 := 0, 0
	for _, v := range position {
		if v%2 == 1 {
			p1++
		} else {
			p2++
		}
	}
	return min(p1, p2)
}
