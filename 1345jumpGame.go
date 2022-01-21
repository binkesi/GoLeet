package goleet

// https://leetcode-cn.com/problems/jump-game-iv/

func minJumps(arr []int) int {
	lenA := len(arr)
	if lenA == 1 {
		return 0
	}
	var graph [][]int = make([][]int, lenA)
	for i := 0; i < lenA; i++ {
		graph[i] = make([]int, 0)
	}
	for i := 0; i < lenA; i++ {
		edges := graph[i]
		for k := i + 1; k < lenA; k++ {
			if len(edges) > 0 {
				break
			}
			if arr[k] == arr[i] {
				for j := 0; j < len(edges); j++ {
					graph[edges[j]] = append(graph[edges[j]], k)
					edges = append(edges, k)
				}
			}
		}
		if i == 0 {
			edges = append(edges, 1)
		} else if i == lenA-1 {
			edges = append(edges, lenA-2)
		} else {
			edges = append(edges, i-1)
			edges = append(edges, i+1)
		}
	}
	minPath := bfs(graph, 0, lenA-1)
	return minPath
}

func bfs(graph [][]int, start, end int) int {
	lenG := len(graph)
	var path int = 0
	var visited []bool = make([]bool, lenG)
	visited[start] = true
	var queue []int = make([]int, 0)
	queue = append(queue, start)
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
	}
}
