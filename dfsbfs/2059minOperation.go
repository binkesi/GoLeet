package dfsbfs

// https://leetcode-cn.com/problems/minimum-operations-to-convert-number/

func minimumOperations(nums []int, start int, goal int) int {
	visited := make([]int, 1001)
	type pairN struct {
		num  int
		step int
	}
	queue := []pairN{{start, 0}}
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		num, step := p.num, p.step
		if num == goal {
			return step
		}
		for _, v := range nums {
			if num+v == goal || num-v == goal || num^v == goal {
				return step + 1
			}
			if num+v >= 0 && num+v <= 1000 && visited[num+v] == 0 {
				visited[num+v] = 1
				queue = append(queue, pairN{num + v, step + 1})
			}
			if num-v >= 0 && num-v <= 1000 && visited[num-v] == 0 {
				visited[num-v] = 1
				queue = append(queue, pairN{num - v, step + 1})
			}
			if num^v >= 0 && num^v <= 1000 && visited[num^v] == 0 {
				visited[num^v] = 1
				queue = append(queue, pairN{num ^ v, step + 1})
			}
		}
	}
	return -1
}
