package weekgame

// https://leetcode-cn.com/contest/weekly-contest-197/problems/path-with-maximum-probability/

type pairP struct {
	end  int
	prob float64
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make(map[int]map[int]float64, n)
	for i := 0; i < n; i++ {
		graph[i] = make(map[int]float64)
	}
	for i, v := range edges {
		graph[v[0]][v[1]] = succProb[i]
		graph[v[1]][v[0]] = succProb[i]
	}
	var ans float64 = 0.0
	visited := make([]float64, n)
	visited[start] = 1.0
	queue := []pairP{{start, 1.0}}
	for len(queue) != 0 {
		p := queue[0]
		if p.end == end {
			if p.prob > ans {
				ans = p.prob
				continue
			}
		}
		queue = queue[1:]
		for i, v := range graph[p.end] {
			cur := visited[p.end]
			if cur*v > visited[i] {
				visited[i] = cur * v
				queue = append(queue, pairP{i, cur * v})
			}
		}
	}
	return ans
}
