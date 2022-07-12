package mathgame

// https://leetcode.cn/problems/cells-with-odd-values-in-a-matrix/submissions/

func oddCells(m int, n int, indices [][]int) int {
	mval, nval := make([]int, m), make([]int, n)
	modd, meven, nodd, neven := 0, 0, 0, 0
	for _, v := range indices {
		mval[v[0]] = (mval[v[0]] + 1) % 2
		nval[v[1]] = (nval[v[1]] + 1) % 2
	}
	for _, v := range mval {
		if v%2 == 0 {
			meven++
		} else {
			modd++
		}
	}
	for _, v := range nval {
		if v%2 == 0 {
			neven++
		} else {
			nodd++
		}
	}
	return meven*nodd + modd*neven
}
