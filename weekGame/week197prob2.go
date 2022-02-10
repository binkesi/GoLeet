package weekgame

// https://leetcode-cn.com/contest/weekly-contest-197/problems/number-of-substrings-with-only-1s/

func numSub(s string) int {
	lenS := len(s)
	ans, tmp := 0, 0
	for i := 0; i < lenS; i++ {
		if s[i] == '1' && (i == 0 || s[i-1] == '0') {
			tmp = 1
		} else if s[i] == '1' {
			tmp += 1
		} else if s[i] == '0' {
			continue
		}
		if s[i] == '1' && (i == lenS-1 || s[i+1] == '0') {
			ans += ((tmp + 1) * tmp / 2) % (1000000007)
		}
	}
	return ans % 1000000007
}
