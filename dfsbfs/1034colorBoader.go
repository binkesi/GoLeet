package dfsbfs

// https://leetcode-cn.com/problems/coloring-a-border/

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	m, n := len(grid), len(grid[0])
	cur := grid[row][col]
	direct := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	type pairC struct {
		x int
		y int
	}
	queue := []pairC{{row, col}}
	res := []pairC{}
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[row][col] = true
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		x, y := p.x, p.y
		cnt := 0
		for _, v := range direct {
			nx, ny := x+v[0], y+v[1]
			if nx >= 0 && ny >= 0 && nx < m && ny < n {
				if grid[nx][ny] == cur {
					cnt += 1
					if !visited[nx][ny] {
						visited[nx][ny] = true
						queue = append(queue, pairC{nx, ny})
					}
				}
			}
		}
		if cnt != 4 {
			res = append(res, pairC{x, y})
		}
	}
	for _, v := range res {
		grid[v.x][v.y] = color
	}
	return grid
}
