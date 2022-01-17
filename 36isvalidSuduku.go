package goleet

// https://leetcode-cn.com/problems/valid-sudoku/

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		numMap := make(map[byte]struct{})
		for j := 0; j < 9; j++ {
			if _, ok := numMap[board[i][j]]; ok {
				return false
			} else if board[i][j] != '.' {
				numMap[board[i][j]] = struct{}{}
			}
		}
	}
	for j := 0; j < 9; j++ {
		numMap := make(map[byte]struct{})
		for i := 0; i < 9; i++ {
			if _, ok := numMap[board[i][j]]; ok {
				return false
			} else if board[i][j] != '.' {
				numMap[board[i][j]] = struct{}{}
			}
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !judgeBlock(board, i, j) {
				return false
			}
		}
	}
	return true
}

func judgeBlock(board [][]byte, l, c int) bool {
	numMap := make(map[byte]struct{})
	for i := l * 3; i < (l+1)*3; i++ {
		for j := c * 3; j < (c+1)*3; j++ {
			if _, ok := numMap[board[i][j]]; ok {
				return false
			} else if board[i][j] != '.' {
				numMap[board[i][j]] = struct{}{}
			}
		}
	}
	return true
}
