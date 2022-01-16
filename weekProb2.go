package goleet

// https://leetcode-cn.com/contest/weekly-contest-276/problems/minimum-moves-to-reach-target-score/

func minMoves(target int, maxDoubles int) int {
	var dp []int = make([]int, target)
	var tmp []int = make([]int, target)
	tmp[0] = 0
	for i := 0; i < target; i++ {
		dp[i] = i
	}
	for l := 1; l <= maxDoubles; l++ {
		for c := 1; c < target; c++ {
			if c%2 == 1 {
				tmp[c] = min(tmp[c-1], dp[(c+1)/2-1]) + 1
			} else {
				tmp[c] = tmp[c-1] + 1
			}
		}
		dp = tmp
	}
	return dp[target-1]
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
