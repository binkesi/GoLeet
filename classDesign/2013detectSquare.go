package classdesign

// https://leetcode-cn.com/problems/detect-squares/

type DetectSquares struct {
	points map[int]map[int]int
}

func Constructor2() DetectSquares {
	return DetectSquares{points: make(map[int]map[int]int)}
}

func (this *DetectSquares) Add(point []int) {
	if _, ok := this.points[point[0]]; !ok {
		this.points[point[0]] = make(map[int]int)
		this.points[point[0]][point[1]] = 1
	} else if _, ok := this.points[point[0]][point[1]]; !ok {
		this.points[point[0]][point[1]] = 1
	} else {
		this.points[point[0]][point[1]] += 1
	}
}

func (this *DetectSquares) Count(point []int) int {
	cnt := 0
	x, y := point[0], point[1]
	tmpx, tmpy := [][]int{}, [][]int{}
	for k, v := range this.points {
		for i, j := range v {
			if k == x && i != y {
				tmpx = append(tmpx, []int{k, i})
			}
			if i == y && k != x {
				tmpy = append(tmpy, []int{k, j})
			}
		}
	}
	for i := 0; i < len(tmpx); i++ {
		for j := 0; j < len(tmpy); j++ {
			if _, ok := this.points[tmpy[j][0]][tmpx[i][1]]; abs(tmpx[i][1], y) == abs(tmpy[j][0], x) && ok {
				cnt += this.points[x][tmpx[i][1]] * this.points[tmpy[j][0]][y] * this.points[tmpy[j][0]][tmpx[i][1]]
			}
		}
	}
	return cnt
}

func abs(a, b int) int {
	if a >= b {
		return a - b
	} else {
		return b - a
	}
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
