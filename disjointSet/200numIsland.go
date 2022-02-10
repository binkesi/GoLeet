package disjointset

// https://leetcode-cn.com/problems/number-of-islands/

func numIslands(grid [][]byte) (ans int) {
	r, c := len(grid), len(grid[0])
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				ans += 1
				queue := [][2]int{{i, j}}
				grid[i][j] = 'x'
				for len(queue) != 0 {
					p := queue[0]
					queue = queue[1:]
					x, y := p[0], p[1]
					for _, v := range dirs {
						nx, ny := x+v[0], y+v[1]
						if nx >= 0 && ny >= 0 && nx < r && ny < c && grid[nx][ny] == '1' {
							queue = append(queue, [2]int{nx, ny})
							grid[nx][ny] = 'x'
						}
					}
				}
			}
		}
	}
	return ans
}
