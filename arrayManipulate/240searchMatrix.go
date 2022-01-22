package arraymanipulate

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/

func searchMatrixZ(matrix [][]int, target int) bool {
	lenx, leny := len(matrix[0]), len(matrix)
	for x, y := lenx-1, leny-1; x < lenx && x >= 0 && y < leny && y >= 0; {
		if matrix[y][x] == target {
			return true
		} else if matrix[y][x] < target {
			x += 1
		} else {
			y -= 1
		}
	}
	return false
}
