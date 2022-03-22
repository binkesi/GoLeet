package greedy

// https://leetcode-cn.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/

func winnerOfGame(colors string) bool {
	cntA, cntB := 0, 0
	n := len(colors)
	for numA, numB, i := 0, 0, 0; i < n; i++ {
		if colors[i] == 'A' {
			numA++
			numB = 0
		} else if colors[i] == 'B' {
			numB++
			numA = 0
		}
		if numA >= 3 {
			cntA++
		}
		if numB >= 3 {
			cntB++
		}
	}
	return cntA > cntB
}
