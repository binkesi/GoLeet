package dfsbfs

// https://leetcode-cn.com/problems/the-time-when-the-network-becomes-idle/

func networkBecomesIdle(edges [][]int, patience []int) (ans int) {
	edgeMap := make(map[int][]int)
	n := len(patience)
	for _, v := range edges {
		edgeMap[v[0]] = append(edgeMap[v[0]], v[1])
		edgeMap[v[1]] = append(edgeMap[v[1]], v[0])
	}
	visited := make([]int, n)
	que := [][2]int{{0, 0}}
	visited[0] = 1
	for len(que) != 0 {
		p := que[0]
		que = que[1:]
		vertex := p[0]
		step := p[1]
		pat := patience[vertex]
		end := 0
		if pat != 0 {
			end = step*4 - (step*2-1)%pat
		}
		if end > ans {
			ans = end
		}
		for _, v := range edgeMap[vertex] {
			if visited[v] == 0 {
				visited[v] = 1
				que = append(que, [2]int{v, step + 1})
			}
		}
	}
	return
}
