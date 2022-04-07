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
