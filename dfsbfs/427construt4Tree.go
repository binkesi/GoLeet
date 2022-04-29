package dfsbfs

// https://leetcode.cn/problems/construct-quad-tree/submissions/

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j != 0 {
				grid[0][j] += grid[0][j-1]
			} else if j == 0 && i != 0 {
				grid[i][0] += grid[i-1][0]
			} else if i != 0 && j != 0 {
				grid[i][j] = grid[i][j] + grid[i-1][j] + grid[i][j-1] - grid[i-1][j-1]
			}
		}
	}
	var dfs func(sx, sy, ex, ey int) *Node
	dfs = func(sx, sy, ex, ey int) *Node {
		node := Node{}
		preSum := (ex - sx + 1) * (ex - sx + 1)
		sum := grid[ex][ey]
		if sx == 0 && sy != 0 {
			sum -= grid[ex][sy-1]
		} else if sx != 0 && sy == 0 {
			sum -= grid[sx-1][ey]
		} else if sx != 0 && sy != 0 {
			sum = sum - grid[ex][sy-1] - grid[sx-1][ey] + grid[sx-1][sy-1]
		}
		if sum == 0 || sum == preSum {
			node.IsLeaf = true
			node.Val = (sum != 0)
			node.TopLeft = nil
			node.TopRight = nil
			node.BottomLeft = nil
			node.BottomRight = nil
			return &node
		} else {
			node.IsLeaf = false
			node.Val = true
			node.TopLeft = dfs(sx, sy, sx+(ex-sx)/2, sy+(ey-sy)/2)
			node.TopRight = dfs(sx, sy+(ey-sy)/2+1, sx+(ex-sx)/2, ey)
			node.BottomLeft = dfs(sx+(ex-sx)/2+1, sy, ex, sy+(ey-sy)/2)
			node.BottomRight = dfs(sx+(ex-sx)/2+1, sy+(ey-sy)/2+1, ex, ey)
		}
		return &node
	}
	return dfs(0, 0, n-1, n-1)
}
