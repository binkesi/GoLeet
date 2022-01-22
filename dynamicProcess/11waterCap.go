// https://leetcode-cn.com/problems/container-with-most-water/
package goleet

import (
	"math"
)

func maxArea(height []int) int {
	length := len(height)
	p1, p2 := 0, length-1
	var maxa int = 0
	for p2 > p1 {
		area := int(math.Min(float64(height[p1]), float64(height[p2]))) * (p2 - p1)
		if maxa < area {
			maxa = area
		}
		if height[p1] < height[p2] {
			p1 += 1
		} else {
			p2 -= 1
		}
	}
	return maxa
}
