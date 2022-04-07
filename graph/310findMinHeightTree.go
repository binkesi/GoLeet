package graph

// https://leetcode-cn.com/problems/minimum-height-trees/

func findMinHeightTrees(n int, edges [][]int) (ans []int) {
	eGraph := make(map[int][]int)
	minDepth := n - 1
	for i := 0; i < n; i++ {
		eGraph[i] = make([]int, 0)
	}
	for _, e := range edges {
		eGraph[e[0]] = append(eGraph[e[0]], e[1])
		eGraph[e[1]] = append(eGraph[e[1]], e[0])
	}
	var dfs func(par, curDep int, maxDep *int, visited []int)
	dfs = func(par, curDep int, maxDep *int, visited []int) {
		visited[par] = 1
		for _, node := range eGraph[par] {
			if visited[node] == 0 {
				dep := curDep + 1
				dfs(node, dep, maxDep, visited)
			} else {
				if curDep > *maxDep {
					*maxDep = curDep
				}
				continue
			}
		}
	}
	for i := 0; i < n; i++ {
		thisDep := 0
		ptr := &thisDep
		visited := make([]int, n)
		dfs(i, 0, ptr, visited)
		if *ptr < minDepth {
			minDepth = *ptr
			ans = make([]int, 0)
			ans = append(ans, i)
		} else if *ptr == minDepth {
			ans = append(ans, i)
		}
	}
	return
}

func findMinHeightTrees2(n int, edges [][]int) []int {
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = []int{}
	}
	// 建图
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// 往下的最大高度
	down1 := make([]int, n)
	// 往下的次大高度
	down2 := make([]int, n)
	// 往上的最大高度
	up := make([]int, n)
	// 记录down1[cur]由哪个孩子节点转移
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = -1
	}
	// 搜索cur往下的最大高度
	var dfs1 func(cur int, fa int) int
	dfs1 = func(cur, fa int) int {
		for _, ne := range adj[cur] {
			if ne == fa {
				continue
			}
			// cur的孩子节点：ne到叶子节点的高度 + 1就是cur到叶子节点的高度
			sub := dfs1(ne, cur) + 1
			if sub > down1[cur] {
				// 最大的成为次大的
				down2[cur] = down1[cur]
				down1[cur] = sub
				p[cur] = ne
			} else if sub > down2[cur] {
				down2[cur] = sub
			}

		}
		return down1[cur]

	}

	// 搜索cur的孩子节点的往上的最大高度
	var dfs2 func(cur int, fa int)
	dfs2 = func(cur, fa int) {
		for _, ne := range adj[cur] {
			if ne == fa {
				continue
			}
			// 先考虑up[ne]由 down[cur]转移过来，这种情况下up[ne]可能从down1[cur]转移或从down2[cur]转移
			if p[cur] != ne {
				up[ne] = max(up[ne], down1[cur]+1)
			} else {
				up[ne] = max(up[ne], down2[cur]+1)

			}
			// 再考虑up[ne] 从 up[cur]转移过来
			up[ne] = max(up[ne], up[cur]+1)
			dfs2(ne, cur)
		}

	}

	dfs1(0, -1)
	dfs2(0, -1)
	//fmt.Println(up, down1)
	res := []int{}
	min := n + 1
	for i := 0; i < n; i++ {
		// 考虑以i为根节点的高度
		h := max(down1[i], up[i])
		if h < min {
			min = h
			res = res[:0]
			res = append(res, i)
		} else if h == min {
			res = append(res, i)
		}

	}
	return res

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
