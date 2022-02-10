package backtracking

import "fmt"

// https://leetcode-cn.com/problems/path-with-maximum-gold/

func getMaximumGold(grid [][]int) (ans int) {
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] != 0 {
				gh := grid
				cur := 0
				fmt.Println(i, j)
				dfs(gh, i, j, cur, &ans)
			}
		}
	}
	return ans
}

func dfs(matrix [][]int, x, y, cur int, res *int) {
	m, n := len(matrix), len(matrix[0])
	cur += matrix[x][y]
	if cur > *res {
		*res = cur
	}
	matrix[x][y] = 0
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, v := range dirs {
		nx, ny := x+v[0], y+v[1]
		if nx >= 0 && ny >= 0 && nx < m && ny < n && matrix[nx][ny] != 0 {
			dfs(matrix, nx, ny, cur, res)
		}
	}
}
