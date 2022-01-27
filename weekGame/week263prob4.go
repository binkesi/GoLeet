package weekgame

import "math"

// https://leetcode-cn.com/contest/weekly-contest-263/problems/second-minimum-time-to-reach-destination/

func secondMinimum(n int, edges [][]int, time int, change int) int {
	graphE := make([][]int, n+1)
	for _, v := range edges {
		graphE[v[0]] = append(graphE[v[0]], v[1])
		graphE[v[1]] = append(graphE[v[1]], v[0])
	}
	type pairE struct {
		p int
		d int
	}
	dist := make([][2]int, n+1)
	dist[1][1] = math.MaxInt32
	for i := 2; i <= n; i++ {
		dist[i] = [2]int{math.MaxInt32, math.MaxInt32}
	}
	queue := []pairE{{1, 0}}
	for dist[n][1] == math.MaxInt32 {
		q := queue[0]
		queue = queue[1:]
		for _, v := range graphE[q.p] {
			curD := q.d + 1
			if dist[v][0] > curD {
				dist[v][0] = curD
				queue = append(queue, pairE{v, curD})
			} else if dist[v][0] < curD && dist[v][1] > curD {
				dist[v][1] = curD
				queue = append(queue, pairE{v, curD})
			}
		}
	}
	step := dist[n][1]
	ans, tmp := 0, 0
	for i := 1; i < step; i++ {
		ans += time
		tmp += time
		if (tmp/change)%2 == 1 {
			ans += (change - tmp%change)
			tmp = 0
		}
	}
	ans += time
	return ans
}
