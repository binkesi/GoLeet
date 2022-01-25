package graph

// https://leetcode-cn.com/problems/surrounded-regions/

func solve2(board [][]byte) {
	if board == nil || len(board) == 0 {
		return
	}
	row, col := len(board), len(board[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			edgeEle := (i == 0 || j == 0 || i == row-1 || j == col-1)
			if edgeEle && board[i][j] == 'O' {
				dfs(board, i, j)
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, x, y int) {
	if x < 0 || y < 0 || x > len(board)-1 || y > len(board[0])-1 || board[x][y] == 'X' || board[x][y] == '#' {
		return
	}
	board[x][y] = '#'
	dfs(board, x-1, y)
	dfs(board, x, y-1)
	dfs(board, x+1, y)
	dfs(board, x, y+1)
}
