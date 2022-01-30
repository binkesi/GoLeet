package dfsbfs

// https://leetcode-cn.com/problems/map-of-highest-peak/

func highestPeak2(isWater [][]int) [][]int {
	row, col := len(isWater), len(isWater[0])
	wats := [][]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if isWater[i][j] == 1 {
				wats = append(wats, []int{i, j})
				isWater[i][j] = 0
			} else {
				isWater[i][j] = -1
			}
		}
	}
	dird := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for h := 1; len(wats) > 0; h++ {
		for size := len(wats); size > 0; size-- {
			p := wats[0]
			wats = wats[1:]
			x, y := p[0], p[1]
			for _, v := range dird {
				nx, ny := x+v[0], y+v[1]
				if nx >= 0 && ny >= 0 && nx < row && ny < col && isWater[nx][ny] == -1 {
					isWater[nx][ny] = h
					wats = append(wats, []int{nx, ny})
				}
			}
		}
	}
	return isWater
}
