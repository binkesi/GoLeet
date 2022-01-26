package graph

import "math"

// https://leetcode-cn.com/problems/second-minimum-time-to-reach-destination/

func secondMinimum2(n int, edges [][]int, time int, change int) int {
	pointGraph := make(map[int][]int)
	for _, v := range edges {
		pointGraph[v[0]] = append(pointGraph[v[0]], v[1])
		pointGraph[v[1]] = append(pointGraph[v[1]], v[0])
	}
	dist := make([][2]int, n+1)
	dist[1][1] = math.MaxInt32
	for i := 2; i <= n; i++ {
		dist[i] = [2]int{math.MaxInt32, math.MaxInt32}
	}
	type ppair struct {
		point    int
		distance int
	}
	queue := []ppair{{1, 0}}
	for dist[n][1] == math.MaxInt32 {
		q := queue[0]
		queue = queue[1:]
		d := q.distance + 1
		for _, v := range pointGraph[q.point] {
			if d < dist[v][0] {
				dist[v][0] = d
				queue = append(queue, ppair{v, d})
			}
			if d > dist[v][0] && d < dist[v][1] {
				dist[v][1] = d
				queue = append(queue, ppair{v, d})
			}
		}
	}
	length := dist[n][1]
	ans, tmp := 0, 0
	for i := 1; i < length; i++ {
		tmp += time
		ans += time
		if (tmp/change)%2 == 1 {
			ans += (change - tmp%change)
			tmp = 0
		}
	}
	ans += time
	return ans
}
