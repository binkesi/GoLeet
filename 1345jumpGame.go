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
		if i == 0 {
			graph[i] = append(graph[i], 1)
		} else if i == lenA-1 {
			graph[i] = append(graph[i], lenA-2)
		} else {
			graph[i] = append(graph[i], i-1)
			graph[i] = append(graph[i], i+1)
		}
	}
	for i := 0; i < lenA; i++ {
		for k := i + 1; k < lenA; k++ {
			if arr[k] == arr[i] {
				graph[i] = append(graph[i], k)
				graph[k] = append(graph[k], i)
			}
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
		lenQ := len(queue)
		for i := 0; i < lenQ; i++ {
			cur := queue[i]
			if cur == end {
				return path
			}
			for _, v := range graph[cur] {
				if visited[v] == false {
					visited[v] = true
					queue = append(queue, v)
				}
			}
		}
		path += 1
		queue = queue[lenQ:]
	}
	return path
}
