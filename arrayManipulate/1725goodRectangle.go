package arraymanipulate

// https://leetcode-cn.com/problems/number-of-rectangles-that-can-form-the-largest-square/

func countGoodRectangles(rectangles [][]int) int {
	ans, large := 0, 0
	for _, v := range rectangles {
		tmp := min(v[0], v[1])
		if tmp > large {
			large = tmp
			ans = 1
		} else if tmp == large {
			ans += 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
