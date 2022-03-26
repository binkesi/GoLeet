package graph

// https://leetcode-cn.com/problems/rotate-image/

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j <= n-i-2; j++ {
			tVal := matrix[i][j]
			for x, y := i, j; ; {
				nx := y
				ny := n - 1 - x
				tmp := matrix[nx][ny]
				matrix[nx][ny] = tVal
				tVal = tmp
				if nx == i && ny == j {
					break
				} else {
					x, y = nx, ny
				}
			}
		}
	}
}
