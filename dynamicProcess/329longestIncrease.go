package dynamicprocess

// https://leetcode-cn.com/problems/longest-increasing-path-in-a-matrix/

func longestIncreasingPath(matrix [][]int) (ans int) {
	r, c := len(matrix), len(matrix[0])
	var bfs func(int, int, [][]int, *int)
	dirs := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	bfs = func(i, j int, graph [][]int, res *int) {
		r, c := len(graph), len(graph[0])
		queue := [][3]int{}
		queue = append(queue, [3]int{i, j, 1})
		for len(queue) != 0 {
			p := queue[0]
			x, y, step := p[0], p[1], p[2]
			queue = queue[1:]
			for _, v := range dirs {
				nx, ny := x+v[0], y+v[1]
				if nx >= 0 && ny >= 0 && nx < r && ny < c && graph[nx][ny] > graph[x][y] {
					queue = append(queue, [3]int{nx, ny, step + 1})
				}
			}
			if len(queue) == 0 {
				if *res < step {
					*res = step
				}
			}
		}
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for ind, v := range dirs {
				nx, ny := i+v[0], i+v[1]
				if nx >= 0 && ny >= 0 && nx < r && ny < r {
					if matrix[nx][ny] < matrix[i][j] {
						break
					}
				}
				if ind == 3 {
					bfs(i, j, matrix, &ans)
				}
			}
		}
	}
	return
}
