package classdesign

import "math/rand"

// https://leetcode.cn/problems/generate-random-point-in-a-circle/

type Solution struct {
	radius   float64
	x_center float64
	y_center float64
}

func Constructor5(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius: radius, x_center: x_center, y_center: y_center}
}

func (this *Solution) RandPoint() []float64 {
	r := this.radius
	for {
		x := r*rand.Float64()*2 - r
		y := r*rand.Float64()*2 - r
		if x*x+y*y <= r*r {
			return []float64{x + this.x_center, y + this.y_center}
		}
	}
}
