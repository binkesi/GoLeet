package dfsbfs

import "math"

// https://leetcode-cn.com/problems/map-of-highest-peak/

func highestPeak(isWater [][]int) [][]int {
	row, col := len(isWater), len(isWater[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if isWater[i][j] == 0 {
				isWater[i][j] = math.MaxInt32
			} else {
				isWater[i][j] = 0
			}
		}
	}
	var bfsWater func(arr [][]int, i, j int)
	type pairw struct {
		x int
		y int
	}
	bfsWater = func(arr [][]int, i, j int) {
		r, c := len(arr), len(arr[0])
		queue := []pairw{{i, j}}
		for len(queue) != 0 {
			p := queue[0]
			x, y := p.x, p.y
			queue = queue[1:]
			if x-1 >= 0 && arr[x-1][y] != 0 && arr[x][y]+1 < arr[x-1][y] {
				arr[x-1][y] = arr[x][y] + 1
				queue = append(queue, pairw{x - 1, y})
			}
			if x+1 < r && arr[x+1][y] != 0 && arr[x][y]+1 < arr[x+1][y] {
				arr[x+1][y] = arr[x][y] + 1
				queue = append(queue, pairw{x + 1, y})
			}
			if y-1 >= 0 && arr[x][y-1] != 0 && arr[x][y]+1 < arr[x][y-1] {
				arr[x][y-1] = arr[x][y] + 1
				queue = append(queue, pairw{x, y - 1})
			}
			if y+1 < c && arr[x][y+1] != 0 && arr[x][y]+1 < arr[x][y+1] {
				arr[x][y+1] = arr[x][y] + 1
				queue = append(queue, pairw{x, y + 1})
			}
		}
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if isWater[i][j] == 0 {
				bfsWater(isWater, i, j)
			}
		}
	}
	return isWater
}
