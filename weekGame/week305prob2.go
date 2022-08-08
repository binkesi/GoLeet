package weekgame

// https://leetcode.cn/contest/weekly-contest-305/problems/reachable-nodes-with-restrictions/

func reachableNodes(n int, edges [][]int, restricted []int) (ans int) {
	restNode := make(map[int]struct{})
	for _, v := range restricted {
		restNode[v] = struct{}{}
	}
	allEdges := make(map[int]map[int]struct{})
	for _, v := range edges {
		if _, ok := allEdges[v[0]]; ok {
			allEdges[v[0]][v[1]] = struct{}{}
		} else {
			allEdges[v[0]] = make(map[int]struct{})
			allEdges[v[0]][v[1]] = struct{}{}
		}
		if _, ok := allEdges[v[1]]; ok {
			allEdges[v[1]][v[0]] = struct{}{}
		} else {
			allEdges[v[1]] = make(map[int]struct{})
			allEdges[v[1]][v[0]] = struct{}{}
		}
	}
	reaches := []int{0}
	ans++
	visited := make([]int, n)
	visited[0] = 1
	for len(reaches) != 0 {
		p := reaches[0]
		reaches = reaches[1:]
		for k, _ := range allEdges[p] {
			if _, ok := restNode[k]; !ok {
				if visited[k] != 1 {
					reaches = append(reaches, k)
					ans++
					visited[k] = 1
				}
			}
		}
	}
	return
}
