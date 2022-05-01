package weekgame

// https://leetcode.cn/contest/weekly-contest-291/problems/total-appeal-of-a-string/

func appealSum(s string) int64 {
	ans, sumG, pos := 0, 0, [26]int{}
	for i := range pos {
		pos[i] = -1
	}
	for i, c := range s {
		c -= 'a'
		sumG += i - pos[c]
		pos[c] = i
		ans += sumG
	}
	return int64(ans)
}
