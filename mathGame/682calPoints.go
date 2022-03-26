package mathgame

import "strconv"

// https://leetcode-cn.com/problems/baseball-game/

func calPoints(ops []string) (ans int) {
	points := []int{}
	for _, op := range ops {
		n := len(points)
		switch op[0] {
		case '+':
			ans += points[n-1] + points[n-2]
			points = append(points, points[n-1]+points[n-2])
		case 'D':
			ans += points[n-1] * 2
			points = append(points, 2*points[n-1])
		case 'C':
			ans -= points[n-1]
			points = points[:len(points)-1]
		default:
			pt, _ := strconv.Atoi(op)
			ans += pt
			points = append(points, pt)
		}
	}
	return
}
