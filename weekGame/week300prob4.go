package weekgame

// https://leetcode.cn/problems/number-of-increasing-paths-in-a-grid/

func countPaths(grid [][]int) (ans int) {
	m, n := len(grid), len(grid[0])
	ways := make([][]int, m)
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	const mod int = 1e9 + 7
	for i := range ways {
		ways[i] = make([]int, n)
	}
	wayQueue := [][3]int{}
	var calWay func(i, j int) (ok bool, way int)
	calWay = func(x, y int) (ok bool, way int) {
		if ways[x][y] != 0 {
			return false, -1
		}
		cur := 1
		for k, v := range dirs {
			dx, dy := x+v[0], y+v[1]
			if dx < 0 || dx >= m || dy < 0 || dy >= n {
				if k != 3 {
					continue
				}
			} else if grid[x][y] < grid[dx][dy] {
				if ways[dx][dy] == 0 {
					break
				}
				cur += (ways[dx][dy] % mod)
			}
			if k == 3 {
				return true, cur
			}
		}
		return false, -1
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ok, v := calWay(i, j); !ok {
				continue
			} else {
				ways[i][j] = v
				wayQueue = append(wayQueue, [3]int{i, j, v})
				ans = (ans + v) % mod
			}
		}
	}
	for len(wayQueue) != 0 {
		q := wayQueue[0]
		x, y := q[0], q[1]
		for _, v := range dirs {
			dx, dy := x+v[0], y+v[1]
			if dx < 0 || dx >= m || dy < 0 || dy >= n {
				continue
			}
			if ok, val := calWay(dx, dy); !ok {
				continue
			} else {
				ways[dx][dy] = val
				wayQueue = append(wayQueue, [3]int{dx, dy, val})
				ans = (ans + val) % mod
			}
		}
		if len(wayQueue) == 1 {
			break
		}
		wayQueue = wayQueue[1:]
	}
	return
}
