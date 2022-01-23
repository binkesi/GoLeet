package graph

// https://leetcode-cn.com/problems/surrounded-regions/

type pair struct{ x, y int }

func solve(board [][]byte) {
	row, col := len(board)-2, len(board[0])-2
	visited := make([][]bool, row)
	for i := range visited {
		visited[i] = make([]bool, col)
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if board[i][j] == 'X' || visited[i-1][j-1] {
				continue
			} else {
				bfsA(board, i, j, visited)
			}
		}
	}
}

func bfsA(board [][]byte, row, col int, visited [][]bool) {
	queue := []pair{{row, col}}
	visited[row-1][col-1] = true
	for i := 0; i < len(queue); i++ {
		p := queue[i]
		r, c := p.x, p.y
		if board[r-1][c] == 'O' {
			if row-1 == 0 {
				for _, v := range queue {
					board[v.x][v.y] = 'O'
				}
				return
			} else if !visited[r-2][c-1] {
				visited[r-2][c-1] = true
				queue = append(queue, pair{r - 1, c})
			}
		}
		if board[r][c-1] == 'O' {
			if c-1 == 0 {
				for _, v := range queue {
					board[v.x][v.y] = 'O'
				}
				return
			} else if !visited[r-1][c-2] {
				visited[r-1][c-2] = true
				queue = append(queue, pair{r, c - 1})
			}
		}
		if board[r+1][c] == 'O' {
			if row+1 == len(board)-1 {
				for _, v := range queue {
					board[v.x][v.y] = 'O'
				}
				return
			} else if !visited[r][c-1] {
				visited[r][c-1] = true
				queue = append(queue, pair{r + 1, c})
			}
		}
		if board[r][c+1] == 'O' {
			if col+1 == len(board[0]) {
				for _, v := range queue {
					board[v.x][v.y] = 'O'
				}
				return
			} else if !visited[r-1][c] {
				visited[r-1][c] = true
				queue = append(queue, pair{r, c + 1})
			}
		}
	}
	for _, v := range queue {
		board[v.x][v.y] = 'X'
	}
}
