package dfsbfs

// https://leetcode.cn/problems/minimum-genetic-mutation/

func diffOne(s, t string) (diff bool) {
	for i := range s {
		if s[i] != t[i] {
			if diff {
				return false
			}
			diff = true
		}
	}
	return
}

func minMutation(start, end string, bank []string) int {
	if start == end {
		return 0
	}

	m := len(bank)
	adj := make([][]int, m)
	endIndex := -1
	for i, s := range bank {
		if s == end {
			endIndex = i
		}
		for j := i + 1; j < m; j++ {
			if diffOne(s, bank[j]) {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	if endIndex == -1 {
		return -1
	}

	var q []int
	vis := make([]bool, m)
	for i, s := range bank {
		if diffOne(start, s) {
			q = append(q, i)
			vis[i] = true
		}
	}
	for step := 1; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			if cur == endIndex {
				return step
			}
			for _, nxt := range adj[cur] {
				if !vis[nxt] {
					vis[nxt] = true
					q = append(q, nxt)
				}
			}
		}
	}
	return -1
}
