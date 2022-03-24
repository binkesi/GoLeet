package mathgame

// https://leetcode-cn.com/problems/image-smoother/

func imageSmoother(img [][]int) [][]int {
	dirs := [][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {0, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}
	m, n := len(img), len(img[0])
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			sum, cnt := 0, 0
			for _, v := range dirs {
				nx, ny := x+v[0], y+v[1]
				if nx >= 0 && ny >= 0 && nx < m && ny < n {
					sum += img[nx][ny]
					cnt++
				}
			}
			ans[x][y] = sum / cnt
		}
	}
	return ans
}
