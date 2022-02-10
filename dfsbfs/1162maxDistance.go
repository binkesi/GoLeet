package dfsbfs

// https://leetcode-cn.com/problems/as-far-from-land-as-possible/

func maxDistance(grid [][]int) int {
	n := len(grid)
	ans := 0
	ground := [][3]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				ground = append(ground, [3]int{i, j, 0})
				grid[i][j] = -1
			}
		}
	}
	if len(ground) == 0 || len(ground) == n*n {
		return -1
	}
	direct := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(ground) != 0 {
		p := ground[0]
		ground = ground[1:]
		x, y, step := p[0], p[1], p[2]
		if step > ans {
			ans = step
		}
		for _, v := range direct {
			nx, ny := x+v[0], y+v[1]
			if nx >= 0 && ny >= 0 && nx < n && ny < n && grid[nx][ny] == 0 {
				grid[nx][ny] = step + 1
				ground = append(ground, [3]int{nx, ny, step + 1})
			}
		}
	}
	return ans
}
