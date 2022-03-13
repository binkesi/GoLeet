package weekgame

// https://leetcode-cn.com/contest/weekly-contest-284/problems/count-artifacts-that-can-be-extracted/

func digArtifacts(n int, artifacts [][]int, dig [][]int) (ans int) {
	dMap := make([][]int, n)
	for i := range dMap {
		dMap[i] = make([]int, n)
	}
	for _, v := range dig {
		dMap[v[0]][v[1]] = 1
	}
	for _, v := range artifacts {
		canDig := true
		x1, y1, x2, y2 := v[0], v[1], v[2], v[3]
		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				if dMap[x][y] == 0 {
					canDig = false
				}
			}
		}
		if canDig {
			ans += 1
		}
	}
	return
}
