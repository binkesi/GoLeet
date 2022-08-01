package stack

// https://leetcode.cn/problems/asteroid-collision/submissions/

func asteroidCollision(asteroids []int) []int {
	ans := []int{}
	l := len(asteroids)
	for i := 0; i < l; {
		if len(ans) == 0 {
			ans = append(ans, asteroids[i])
			i++
			continue
		}
		n := len(ans)
		p := ans[n-1]
		if p*asteroids[i] > 0 || p < 0 && asteroids[i] > 0 {
			ans = append(ans, asteroids[i])
			i++
			continue
		} else if abs(asteroids[i]) < p {
			i++
			continue
		} else if abs(asteroids[i]) == p {
			ans = ans[:n-1]
			i++
			continue
		} else {
			ans = ans[:n-1]
			continue
		}
	}
	return ans
}

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
