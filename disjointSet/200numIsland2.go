package disjointset

// https://leetcode-cn.com/problems/number-of-islands/

func numIslands(grid [][]byte) (ans int) {
	r, c := len(grid), len(grid[0])
	dirs := [][2]int{{-1, 0}, {0, -1}}
	uf := NewisUF(grid)
	islandMap := make(map[int]struct{})
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if grid[i][j] == '1' {
				idx := i*c + j
				root := uf.Find(idx)
				islandMap[root] = struct{}{}
				for _, v := range dirs {
					nx, ny := i+v[0], j+v[1]
					ndx := nx*c + ny
					if nx >= 0 && ny >= 0 && nx < r && ny < c && grid[nx][ny] == '1' {
						if uf.Connected(idx, ndx) {
							continue
						}
						uf.Union(idx, ndx)
						root = uf.Find(ndx)
						islandMap[root] = struct{}{}
					}
				}
			}
		}
	}
	ans = len(islandMap)
	return
}

type islandUF struct {
	parent [][]int
}

func NewisUF(grid [][]byte) *islandUF {
	r, c := len(grid), len(grid[0])
	uf := islandUF{}
	uf.parent = make([][]int, r)
	for i := range uf.parent {
		uf.parent[i] = make([]int, c)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			uf.parent[i][j] = i*r + j
		}
	}
	return &uf
}

func (uf *islandUF) Union(i, j int) {
	pi, pj := uf.Find(i), uf.Find(j)
	if pi == pj {
		return
	}
	c := len(uf.parent[0])
	if pi < pj {
		x, y := pj/c, pj%c
		uf.parent[x][y] = pi
	} else {
		x, y := pi/c, pj%c
		uf.parent[x][y] = pj
	}
}

func (uf *islandUF) Connected(i, j int) bool {
	return uf.Find(i) == uf.Find(j)
}

func (uf *islandUF) Find(i int) int {
	root := i
	c := len(uf.parent[0])
	x, y := i/c, i%c
	for uf.parent[x][y] != root {
		root = uf.parent[x][y]
		x, y = root/c, root%c
	}
	for i != root {
		x, y = i/c, i%c
		i, uf.parent[x][y] = uf.parent[x][y], root
	}
	return root
}
