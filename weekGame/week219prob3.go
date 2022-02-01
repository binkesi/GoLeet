package weekgame

// https://leetcode-cn.com/contest/weekly-contest-219/problems/stone-game-vii/

func stoneGameVII(stones []int) int {
	ans := AmoveStone(stones, 0)
	return ans
}

func AmoveStone(stones []int, cur int) int {
	lenS := len(stones)
	if lenS <= 1 {
		return cur
	}
	if BmoveStone(stones[1:], cur) > BmoveStone(stones[0:lenS-1], cur) {
		return BmoveStone(stones[1:], cur)
	} else {
		return BmoveStone(stones[0:lenS-1], cur)
	}
}

func BmoveStone(stones []int, cur int) int {
	lenS := len(stones)
	if lenS == 0 {
		return cur
	}
	if lenS == 1 {
		return cur + stones[0]
	}
	if AmoveStone(stones[1:], cur+stones[0]) > AmoveStone(stones[0:lenS-1], cur+stones[lenS-1]) {
		return AmoveStone(stones[0:lenS-1], cur+stones[lenS-1])
	} else {
		return AmoveStone(stones[1:], cur+stones[0])
	}
}
