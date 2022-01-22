package mathgame

// https://leetcode-cn.com/problems/stone-game-ix/

func stoneGameIX(stones []int) bool {
	lenS := len(stones)
	if lenS <= 1 {
		return false
	}
	n0, n1, n2 := 0, 0, 0
	for i := 0; i < lenS; i++ {
		if stones[i]%3 == 0 {
			n0 += 1
		} else if stones[i]%3 == 1 {
			n1 += 1
		} else {
			n2 += 1
		}
	}
	res := (n0%2 == 0)
	if res {
		return n1 >= 1 && n2 >= 1
	} else {
		return abs(n1, n2) > 2
	}
}

func abs(a, b int) int {
	if a-b >= 0 {
		return a - b
	} else {
		return b - a
	}
}
