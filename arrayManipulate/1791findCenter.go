package arraymanipulate

// https://leetcode-cn.com/problems/find-center-of-star-graph/

func findCenter(edges [][]int) int {
	x, y := edges[0][0], edges[0][1]
	nx, ny := edges[1][0], edges[1][1]
	if x == nx || x == ny {
		return x
	} else {
		return y
	}
}
