package mathgame

// https://leetcode-cn.com/problems/count-of-matches-in-tournament/

func numberOfMatches(n int) int {
	res := 0
	for n != 1 {
		res += n / 2
		n = n/2 + n%2
	}
	return res
}
