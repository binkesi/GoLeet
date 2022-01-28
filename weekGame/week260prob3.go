package weekgame

// https://leetcode-cn.com/problems/check-if-word-can-be-placed-in-crossword/

func placeWordInCrossword(board [][]byte, word string) bool {
	row, col := len(board), len(board[0])
	lenW := len(word)
	for i := 0; i < row; i++ {
		for j := 0; j <= col-lenW; j++ {
			if match(board[i][j:j+lenW], word) || match(board[i][j:j+lenW], reverse(word)) {
				if j == 0 && j+lenW < col && board[i][j+lenW] != ' ' {
					return true
				} else if j+lenW == col && j > 0 && board[i][j-1] != ' ' {
					return true
				} else if board[i][j-1] != ' ' && board[i][j+lenW] != ' ' {
					return true
				}
			}
		}
	}
	for j := 0; j < col; j++ {
		for i := 0; i <= row-lenW; i++ {
			tmp := make([]byte, 0)
			for k := i; k <= i+lenW-1; k++ {
				tmp = append(tmp, board[k][j])
			}
			if match(tmp, word) || match(tmp, reverse(word)) {
				if i == 0 && i+lenW < row && board[i+lenW][j] != ' ' {
					return true
				} else if i+lenW == col && i > 0 && board[i-1][j] != ' ' {
					return true
				} else if board[i-1][j] != ' ' && board[i+lenW][j] != ' ' {
					return true
				}
			}
		}
	}
	return false
}

func match(bytes []byte, word string) bool {
	lenB := len(bytes)
	for i := 0; i < lenB; i++ {
		if bytes[i] == '#' {
			return false
		}
		if bytes[i] == ' ' {
			continue
		} else if bytes[i] != word[i] {
			return false
		}
	}
	return true
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
