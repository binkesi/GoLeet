package weekgame

// https://leetcode-cn.com/contest/weekly-contest-196/problems/count-submatrices-with-all-ones/

func numSubmat(mat [][]int) int {
	m := len(mat)
	if m == 0 {
		return 0
	}
	n := len(mat[0])
	sum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j >= 1 && mat[i][j] == 1 {
				mat[i][j] = mat[i][j-1] + 1
			}
			tmp := mat[i][j]
			for k := i; k >= 0; k-- {
				if tmp > mat[k][j] {
					tmp = mat[k][j]
				}
				sum += tmp
			}
		}
	}
	return sum
}
