package arraymanipulate

import "math"

// https://leetcode-cn.com/problems/lucky-numbers-in-a-matrix/

func luckyNumbers(matrix [][]int) (ans []int) {
	r, c := len(matrix), len(matrix[0])
	for i := 0; i < r; i++ {
		lmin, lind := math.MaxInt32, 0
		for j := 0; j < c; j++ {
			if matrix[i][j] < lmin {
				lmin = matrix[i][j]
				lind = j
			}
		}
		for k := 0; k < r; k++ {
			if matrix[k][lind] > lmin {
				break
			}
			if k == r-1 {
				ans = append(ans, lmin)
			}
		}
	}
	return
}
