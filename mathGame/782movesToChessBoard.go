package mathgame

// https://leetcode.cn/problems/transform-to-chessboard/

func movesToChessboard(board [][]int) (ans int) {
	l := len(board[0])
	num := board[0][0]
	a, b, standA, standB := 0, 1, 0, 1
	for i := 1; i < l; i++ {
		standA <<= 1
		standA += b
		standB <<= 1
		standB += a
		a, b = b, a
	}
	for i := 1; i < l; i++ {
		num <<= 1
		num += board[0][i]
	}
	resA, resB := num^standA, num^standB
	arrA, arrB, arr := make([]int, l), make([]int, l), make([]int, l)
	cntA, cntB := 0, 0
	for i := 0; i < l; i++ {
		arrA[l-1-i] = resA & 1
		arrB[l-1-i] = resB & 1
		if resA&1 == 1 {
			cntA++
		}
		if resB&1 == 1 {
			cntB++
		}
		resA >>= 1
		resB >>= 1
	}
	if cntA > cntB {
		arr = arrA
	} else {
		arr = arrB
	}

}

func isChess(line []int) bool {
	n := len(line)
	cur, cnt := line[0], 0
	for i := 1; i < n; i++ {
		if cur == 1 {
			cnt++
		}
		if line[i] == cur {
			return false
		}
		cur = line[i]
	}
	if cnt == n/2 || (n%2 == 1 && cnt == n/2+1) {
		return true
	}
	return false
}
