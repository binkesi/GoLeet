package disjointset

import "sort"

// https://leetcode-cn.com/problems/number-of-provinces/

func findCircleNum(isConnected [][]int) (ans int) {
	lenI := len(isConnected)
	uf := NewUFcity(isConnected)
	for i := 0; i < lenI; i++ {
		for j := i + 1; j < lenI; j++ {
			if isConnected[i][j] == 1 {
				uf.Union(i, j)
			}
		}
	}
	for i := range uf.parent {
		uf.Find(i)
	}
	sort.Ints(uf.parent)
	ans = 1
	for i := 1; i < lenI; i++ {
		if uf.parent[i] != uf.parent[i-1] {
			ans += 1
		}
	}
	return
}

type UFcity struct {
	parent []int
}

func NewUFcity(cityCon [][]int) *UFcity {
	lenC := len(cityCon)
	uf := UFcity{}
	for i := 0; i < lenC; i++ {
		uf.parent = append(uf.parent, i)
	}
	return &uf
}

func (uf *UFcity) Union(i, j int) {
	pi, pj := uf.Find(i), uf.Find(j)
	if pi == pj {
		return
	}
	if pi < pj {
		uf.parent[pj] = pi
	} else {
		uf.parent[pi] = pj
	}
}

func (uf *UFcity) Find(i int) int {
	root := i
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	for uf.parent[i] != root {
		i, uf.parent[i] = uf.parent[i], root
	}
	return root
}
