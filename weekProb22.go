package goleet

// https://leetcode-cn.com/contest/weekly-contest-276/problems/minimum-moves-to-reach-target-score/

func minMovesSec(target int, maxDoubles int) int {
	var res int = 0
	for target != 1 {
		if target%2 == 1 {
			target -= 1
			res += 1
		} else if maxDoubles > 0 {
			target = target >> 1
			res += 1
			maxDoubles -= 1
		} else {
			res += (target - 1)
			break
		}
	}
	return res
}
