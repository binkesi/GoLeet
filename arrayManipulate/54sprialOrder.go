package arraymanipulate

// https://leetcode-cn.com/problems/spiral-matrix/

func spiralOrder(matrix [][]int) (ans []int) {
	m, n := len(matrix), len(matrix[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for x, y, cnt := 0, 0, 0; ; {
		ans = append(ans, matrix[x][y])
		visited[x][y] = true
		dir := dirs[cnt%4]
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && ny >= 0 && nx < m && ny < n && visited[nx][ny] == false {
			x, y = nx, ny
		} else {
			cnt++
			dir := dirs[cnt%4]
			nx, ny = x+dir[0], y+dir[1]
			if nx >= 0 && ny >= 0 && nx < m && ny < n && visited[nx][ny] == false {
				x, y = nx, ny
			} else {
				break
			}
		}
	}
	return
}
