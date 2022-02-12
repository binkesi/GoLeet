package disjointset

// https://leetcode-cn.com/problems/number-of-enclaves/

type unionFind struct {
	parent []int
	rank   []int
	onEdge []bool
}

func newUnionFind(grid [][]int) unionFind {
	m, n := len(grid), len(grid[0])
	parent := make([]int, m*n)
	rank := make([]int, m*n)
	onEdge := make([]bool, m*n)
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				idx := i*n + j
				parent[idx] = idx
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					onEdge[idx] = true
				}
			}
		}
	}
	return unionFind{parent, rank, onEdge}
}

func (uf unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf unionFind) merge(x, y int) {
	x, y = uf.find(x), uf.find(y)
	if x == y {
		return
	}
	if uf.rank[x] > uf.rank[y] {
		uf.parent[y] = x
		uf.onEdge[x] = uf.onEdge[x] || uf.onEdge[y]
	} else if uf.rank[x] < uf.rank[y] {
		uf.parent[x] = y
		uf.onEdge[y] = uf.onEdge[y] || uf.onEdge[x]
	} else {
		uf.parent[y] = x
		uf.onEdge[x] = uf.onEdge[x] || uf.onEdge[y]
		uf.rank[x]++
	}
}

func numEnclaves(grid [][]int) (ans int) {
	uf := newUnionFind(grid)
	m, n := len(grid), len(grid[0])
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				idx := i*n + j
				if j+1 < n && grid[i][j+1] == 1 {
					uf.merge(idx, idx+1)
				}
				if i+1 < m && grid[i+1][j] == 1 {
					uf.merge(idx, idx+n)
				}
			}
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 1 && !uf.onEdge[uf.find(i*n+j)] {
				ans++
			}
		}
	}
	return
}
