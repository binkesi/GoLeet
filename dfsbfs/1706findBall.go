package dfsbfs

// https://leetcode-cn.com/problems/where-will-the-ball-fall/

func findBall(grid [][]int) (ans []int) {
	c := len(grid[0])
	var dfs func(int, int, [][]int, *[]int)
	dfs = func(x, y int, graph [][]int, res *[]int) {
		r, c := len(graph), len(graph[0])
		if x == r {
			*res = append(*res, y)
			return
		}
		val := graph[x][y]
		if (y == 0 && val == -1) || (y == c-1 && val == 1) {
			*res = append(*res, -1)
			return
		}
		if val == 1 {
			if graph[x][y+1] == -1 {
				*res = append(*res, -1)
				return
			} else {
				dfs(x+1, y+1, graph, res)
			}
		} else if val == -1 {
			if graph[x][y-1] == 1 {
				*res = append(*res, -1)
				return
			} else {
				dfs(x+1, y-1, graph, res)
			}
		}
	}
	for i := 0; i < c; i++ {
		dfs(0, i, grid, &ans)
	}
	return
}
