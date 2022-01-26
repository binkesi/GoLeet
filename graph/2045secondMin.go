package graph

import "fmt"

// https://leetcode-cn.com/problems/second-minimum-time-to-reach-destination/

type ppair struct {
	prev int
	cur  int
}

func secondMinimum(n int, edges [][]int, time int, change int) int {
	path, queue := []int{n}, []ppair{{0, 1}}
	visited := make([]bool, n)
	visited[0] = true
	pointGraph := make(map[int][]int)
	for _, v := range edges {
		pointGraph[v[0]] = append(pointGraph[v[0]], v[1])
		pointGraph[v[1]] = append(pointGraph[v[1]], v[0])
	}
new:
	for i := 0; ; i++ {
		q := queue[i]
		point := q.cur
		for _, v := range pointGraph[point] {
			if !visited[v-1] {
				visited[v-1] = true
				queue = append(queue, ppair{point, v})
			}
			if v == n {
				break new
			}
		}
	}
	prev := queue[len(queue)-1].prev
	for i := len(queue) - 1; i >= 0; i-- {
		a, b := queue[i].cur, queue[i].prev
		if a == prev {
			path = append([]int{a}, path...)
			prev = b
		}
	}
	fmt.Println(path)
	return 0
}
