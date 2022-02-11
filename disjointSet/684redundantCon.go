package disjointset

// https://leetcode-cn.com/problems/redundant-connection/

func findRedundantConnection(edges [][]int) (ans []int) {
	lenE := len(edges)
	for idx := lenE - 1; idx >= 0; idx-- {
		uf := NewEdges(edges)
		for i, v := range edges {
			if i == idx {
				continue
			}
			uf.Union(v[0], v[1])
		}
		for i, _ := range uf.parent {
			_ = uf.Find(i)
		}
		res := 0
		for i := 1; i < len(uf.parent); i++ {
			if uf.parent[i] == i {
				res += 1
			}
		}
		if res == 1 {
			ans = edges[idx]
			return
		}
	}
	return
}

type UFedges struct {
	parent []int
}

func NewEdges(edges [][]int) *UFedges {
	uf := UFedges{}
	lenE := len(edges)
	for i := 0; i <= lenE; i++ {
		uf.parent = append(uf.parent, i)
	}
	return &uf
}

func (uf *UFedges) Union(i, j int) {
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

func (uf *UFedges) Find(i int) int {
	root := i
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	for uf.parent[i] != root {
		i, uf.parent[i] = uf.parent[i], root
	}
	return root
}
