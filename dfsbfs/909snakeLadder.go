package dfsbfs

// https://leetcode-cn.com/problems/snakes-and-ladders/

func snakesAndLadders(board [][]int) (ans int) {
	row := len(board)
	if board[0][0] == 1 {
		return -1
	}
	flat := []int{}
	for i := row - 1; i >= 0; i-- {
		if (row-i)%2 == 1 {
			flat = append(flat, board[i]...)
		} else {
			flat = append(flat, reverse(board[i])...)
		}
	}
	type pairB struct {
		num  int
		step int
	}
	queue := []pairB{{1, 0}}
	visited := make([]bool, row*row)
	visited[0] = true
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		num, step := p.num, p.step
		if num == row*row {
			ans = step
			return
		}
		for i := 1; i <= 6; i++ {
			if num+i > row*row {
				break
			}
			if flat[num+i-1] == -1 && visited[num+i-1] == false {
				visited[num+i-1] = true
				queue = append(queue, pairB{num + i, step + 1})
			} else if flat[num+i-1] != -1 && visited[flat[num+i-1]-1] == false {
				visited[flat[num+i-1]-1] = true
				queue = append(queue, pairB{flat[num+i-1], step + 1})
			}
		}
	}
	return -1
}

func reverse(a []int) []int {
	lenA := len(a)
	for i := 0; i <= lenA/2-1; i++ {
		a[i], a[lenA-i-1] = a[lenA-i-1], a[i]
	}
	return a
}
