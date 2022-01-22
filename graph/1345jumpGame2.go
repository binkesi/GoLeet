package graph

// https://leetcode-cn.com/problems/jump-game-iv/

func minJumps2(arr []int) int {
	lenA := len(arr)
	vgraph := make(map[int][]int, 0)
	for i, v := range arr {
		vgraph[v] = append(vgraph[v], i)
	}
	visited := map[int]bool{0: true}
	type pair struct{ idx, step int }
	q := []pair{{}}
	for {
		p := q[0]
		q = q[1:]
		i, step := p.idx, p.step
		if i == lenA-1 {
			return step
		}
		for _, j := range vgraph[arr[i]] {
			if !visited[j] {
				visited[j] = true
				q = append(q, pair{j, step + 1})
			}
		}
		delete(vgraph, arr[i])
		if !visited[i+1] {
			visited[i+1] = true
			q = append(q, pair{i + 1, step + 1})
		}
		if i >= 1 && !visited[i-1] {
			visited[i-1] = true
			q = append(q, pair{i - 1, step + 1})
		}
	}
}
