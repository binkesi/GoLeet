package disjointset

// https://leetcode-cn.com/problems/evaluate-division/

func calcEquation(equations [][]string, values []float64, queries [][]string) (ans []float64) {
	res := make(map[string]map[string]float64)
	for i, v := range equations {
		if _, ok := res[v[0]]; !ok {
			res[v[0]] = make(map[string]float64)
		}
		res[v[0]][v[1]] = values[i]
		if _, ok := res[v[1]]; !ok {
			res[v[1]] = make(map[string]float64)
		}
		res[v[1]][v[0]] = 1 / values[i]
	}
	for _, v := range queries {
		s, e := v[0], v[1]
		if _, ok := res[s]; !ok {
			ans = append(ans, float64(-1))
			continue
		}
		queue := []string{s}
		val := []float64{1}
		visited := make(map[string]struct{})
		visited[s] = struct{}{}
		for len(queue) != 0 {
			p := queue[0]
			v := val[0]
			queue = queue[1:]
			val = val[1:]
			if p == e {
				ans = append(ans, v)
				break
			} else if _, ok := res[p]; ok {
				for k, vs := range res[p] {
					if _, ok := visited[k]; !ok {
						visited[k] = struct{}{}
						queue = append(queue, k)
						val = append(val, v*vs)
					}
				}
			}
			if len(queue) == 0 {
				ans = append(ans, float64(-1))
			}
		}
	}
	return
}
