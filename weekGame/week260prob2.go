package weekgame

// https://leetcode-cn.com/contest/weekly-contest-260/problems/grid-game/

func gridGame(grid [][]int) int64 {
	col := len(grid[0])
	if col <= 1 {
		return 0
	}
	sum0, sum1 := make([]int, col), make([]int, col)
	sum0[0], sum1[0] = grid[0][0], grid[1][0]
	for i := 1; i < col; i++ {
		sum0[i], sum1[i] = grid[0][i]+sum0[i-1], grid[1][i]+sum1[i-1]
	}
	ans := sum0[col-1] - sum0[0]
	if sum1[col-2] < ans {
		ans = sum1[col-2]
	}
	for i := 1; i < col-1; i++ {
		tmp := maxNum(sum0[col-1]-sum0[i], sum1[i-1])
		if tmp < ans {
			ans = tmp
		}
	}
	return int64(ans)
}

func maxNum(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
