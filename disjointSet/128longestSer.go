package disjointset

// https://leetcode-cn.com/problems/longest-consecutive-sequence/

func longestConsecutive(nums []int) (ans int) {
	uf := NewUF(nums)
	for i := 0; i < len(nums); i++ {
		num, tmp := nums[i], nums[i]
		for {
			if _, ok := uf.Find(num - 1); !ok {
				break
			}
			if !uf.Connected(num, num-1) {
				uf.Union(num, num-1)
				num = num - 1
			} else {
				num, _ = uf.Find(num)
				break
			}
		}
		ans = max(ans, tmp-num+1)
	}
	return
}

type UF struct {
	parent map[int]int
}

func NewUF(nums []int) *UF {
	uf := UF{}
	uf.parent = make(map[int]int)
	for _, v := range nums {
		uf.parent[v] = v
	}
	return &uf
}

func (uf *UF) Union(i, j int) {
	pi, _ := uf.Find(i)
	pj, _ := uf.Find(j)
	if pi == pj {
		return
	}
	if pi < pj {
		uf.parent[pj] = pi
	} else {
		uf.parent[pi] = pj
	}
}

func (uf *UF) Connected(i, j int) bool {
	pi, _ := uf.Find(i)
	pj, _ := uf.Find(j)
	return pi == pj
}

func (uf *UF) Find(i int) (int, bool) {
	if _, ok := uf.parent[i]; !ok {
		return -1, false
	}
	root := i
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	for uf.parent[i] != root {
		i, uf.parent[i] = uf.parent[i], root
	}
	return root, true
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
