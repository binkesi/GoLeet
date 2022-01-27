package graph

// https://leetcode-cn.com/problems/surrounded-regions/

type pair struct{ x, y int }

func solve(board [][]byte) {
	row, col := len(board), len(board[0])
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if board[i][j] == 'X' || visited[i][j] {
				continue
			} else {
				bfsA(board, i, j, visited)
			}
		}
	}
}

func bfsA(board [][]byte, row, col int, visited [][]bool) {
	alone := true
	queue := []pair{{row, col}}
	visited[row][col] = true
	for i := 0; i < len(queue); i++ {
		p := queue[i]
		r, c := p.x, p.y
		if r == 0 || r == len(board)-1 || c == 0 || c == len(board[0])-1 {
			alone = false
		}
		if r-1 >= 0 && board[r-1][c] == 'O' && !visited[r-1][c] {
			if r-1 == 0 {
				alone = false
			}
			visited[r-1][c] = true
			queue = append(queue, pair{r - 1, c})
		}
		if c-1 >= 0 && board[r][c-1] == 'O' && !visited[r][c-1] {
			if c-1 == 0 {
				alone = false
			}
			visited[r][c-1] = true
			queue = append(queue, pair{r, c - 1})
		}
		if r+1 <= len(board)-1 && board[r+1][c] == 'O' && !visited[r+1][c] {
			if r+1 == len(board)-1 {
				alone = false
			}
			visited[r+1][c] = true
			queue = append(queue, pair{r + 1, c})
		}
		if c+1 <= len(board[0])-1 && board[r][c+1] == 'O' && !visited[r][c+1] {
			if c+1 == len(board[0])-1 {
				alone = false
			}
			visited[r][c+1] = true
			queue = append(queue, pair{r, c + 1})
		}
	}
	if !alone {
		return
	}
	for _, v := range queue {
		board[v.x][v.y] = 'X'
	}
}
