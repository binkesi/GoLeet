package mathgame

// https://leetcode.cn/problems/projection-area-of-3d-shapes/submissions/

func projectionArea(grid [][]int) (ans int) {
	xy, yz, zx := 0, 0, 0
	for _, v := range grid {
		tmp := 0
		for i, val := range v {
			if val != 0 {
				xy++
			}
			if val > tmp {
				tmp = val
			}
			if i == len(v)-1 {
				yz += tmp
			}
		}
	}
	n := len(grid[0])
	for i := 0; i < n; i++ {
		tmp := 0
		for j, v := range grid {
			if v[i] > tmp {
				tmp = v[i]
			}
			if j == len(grid)-1 {
				zx += tmp
			}
		}
	}
	ans = xy + yz + zx
	return
}
