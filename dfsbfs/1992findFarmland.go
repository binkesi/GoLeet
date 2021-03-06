package dfsbfs

// https://leetcode.cn/problems/find-all-groups-of-farmland/submissions/

func findFarmland(land [][]int) (ans [][]int) {
	r, c := len(land), len(land[0])
	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	visited := make([][]int, r)
	for i := range visited {
		visited[i] = make([]int, c)
	}
	var travel func(x, y int)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if visited[i][j] == 1 || land[i][j] == 0 {
				continue
			}
			visited[i][j] = 1
			minx, miny, maxx, maxy := i, j, i, j
			travel = func(x, y int) {
				visited[x][y] = 1
				if x < minx {
					minx = x
				}
				if y < miny {
					miny = y
				}
				if x > maxx {
					maxx = x
				}
				if y > maxy {
					maxy = y
				}
				for _, v := range dirs {
					dx, dy := x+v[0], y+v[1]
					if dx >= 0 && dx < r && dy >= 0 && dy < c && visited[dx][dy] == 0 && land[dx][dy] == 1 {
						travel(dx, dy)
					}
				}
			}
			travel(i, j)
			ans = append(ans, []int{minx, miny, maxx, maxy})
		}
	}
	return
}

func findFarmland2(land [][]int) (ans [][]int) {
	r, c := len(land), len(land[0])
	for i, row := range land {
		for j, sqr := range row {
			if sqr == 0 || (j > 0 && land[i][j-1] == 1) || (i > 0 && land[i-1][j] == 1) {
				continue
			}
			mi, mj := i, j
			for k := i + 1; k < r; k++ {
				if land[k][j] == 1 {
					mi++
				} else {
					break
				}
			}
			for k := j + 1; k < c; k++ {
				if land[i][k] == 1 {
					mj++
				} else {
					break
				}
			}
			ans = append(ans, []int{i, j, mi, mj})
		}
	}
	return
}
