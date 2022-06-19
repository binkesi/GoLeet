package arraymanipulate

// https://leetcode.cn/problems/diagonal-traverse/submissions/

func findDiagonalOrder(mat [][]int) (ans []int) {
	m, n := len(mat), len(mat[0])
	up := m + n - 2
	for sum, direct := 0, 1; sum <= up; sum++ {
		if direct == 1 {
			for row := m - 1; row >= 0; row-- {
				col := sum - row
				if col < 0 || col > n-1 {
					continue
				}
				ans = append(ans, mat[row][col])

			}
		} else {
			for row := 0; row < m; row++ {
				col := sum - row
				if col > n-1 || col < 0 {
					continue
				}
				ans = append(ans, mat[row][col])
			}
		}
		direct *= -1
	}
	return
}
