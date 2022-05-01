package weekgame

// https://leetcode.cn/contest/weekly-contest-291/problems/remove-digit-from-number-to-maximize-result/

func removeDigit(number string, digit byte) (ans string) {
	n := len(number)
	ind := 0
	for i := 0; i < n; i++ {
		if number[i] == digit {
			if i == n-1 {
				ind = i
			} else if number[i+1] > number[i] {
				ind = i
				break
			}
			ind = i
		}
	}
	if ind != n-1 {
		ans = number[0:ind] + number[ind+1:]
	}
	if ind == n-1 {
		ans = number[0:ind]
	}
	return
}
