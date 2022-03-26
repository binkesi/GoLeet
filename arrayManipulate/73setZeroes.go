package arraymanipulate

// https://leetcode-cn.com/problems/set-matrix-zeroes/

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col0 := false
	for _, v := range matrix {
		if v[0] == 0 {
			col0 = true
		}
		for j := 1; j < n; j++ {
			if v[j] == 0 {
				v[0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
