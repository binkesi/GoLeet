package dfsbfs

// https://leetcode-cn.com/problems/map-of-highest-peak/

func highestPeak2(isWater [][]int) [][]int {
	wats := [][]int{}
	m, n := len(isWater), len(isWater[0])
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if isWater[x][y] == 1 {
				isWater[x][y] = 0
				wats = append(wats, []int{x, y})
			} else {
				isWater[x][y] = -1
			}
		}
	}
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(wats) > 0 {
		q := wats[0]
		wats = wats[1:]
		x, y := q[0], q[1]
		val := isWater[x][y]
		for _, v := range dirs {
			nx, ny := x+v[0], y+v[1]
			if nx >= 0 && ny >= 0 && nx < m && ny < n && isWater[nx][ny] == -1 {
				isWater[nx][ny] = val + 1
				wats = append(wats, []int{nx, ny})
			}
		}
	}
	return isWater
}
