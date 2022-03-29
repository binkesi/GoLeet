package twopointer

// https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam/

func maxConsecutiveAnswers(answerKey string, k int) int {
	n := len(answerKey)
	var getCnt func(a byte) int
	getCnt = func(a byte) int {
		res := 0
		for cnt, i, j := 0, 0, 0; j < n; j++ {
			if answerKey[j] == a {
				cnt++
			}
			for cnt > k {
				if answerKey[i] == a {
					cnt--
				}
				i++
			}
			res = max(res, j-i+1)
		}
		return res
	}
	ans := max(getCnt('T'), getCnt('F'))
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
