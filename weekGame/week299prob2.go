package weekgame

// https://leetcode.cn/contest/weekly-contest-299/problems/count-number-of-ways-to-place-houses/

func countHousePlacements(n int) int {
	a, b := 1, 2
	if n == 1 {
		return 4
	}
	for i := 2; i <= n; i++ {
		tmp := a
		a = b
		b += tmp
		b %= 1000000007
	}
	return (b * b) % 1000000007
}
