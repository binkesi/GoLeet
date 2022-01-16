package goleet

// https://leetcode-cn.com/contest/weekly-contest-276/problems/solving-questions-with-brainpower/

func mostPoints(questions [][]int) int64 {
	lenQ := len(questions)
	if lenQ == 1 {
		return int64(questions[0][0])
	}
	var dp []int = make([]int, lenQ)
	dp[lenQ-1] = questions[lenQ-1][0]
	for i := lenQ - 2; i >= 0; i-- {
		braini := questions[i][1]
		scorei := questions[i][0]
		if i+braini <= lenQ-2 {
			dp[i] = scorei + dp[i+braini+1]
		} else {
			dp[i] = scorei
		}
		if dp[i] < dp[i+1] {
			dp[i] = dp[i+1]
		}
	}
	return int64(dp[0])
}
