package weekgame

// https://leetcode-cn.com/contest/weekly-contest-278/problems/find-substring-with-given-hash-value/

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	lenS := len(s)
	res := 0
	for i := 0; i < lenS-k+1; i++ {
		substr := s[i : i+k]
		if res == 0 {
			for j := 0; j < k; j++ {
				val := (power1(power, j, modulo) % modulo) * int(substr[j]-'a'+1)
				res += (val % modulo)
			}
		} else {
			res = (res*power - (int(s[i-1] - 'a' + 1)) + int(s[i+k-1]-'a'+1)*(power1(power, k-1, modulo)%modulo)) % modulo
		}
		if res%modulo == hashValue {
			return substr
		}
	}
	return s
}

func power1(a, n, mod int) int {
	ans := 1
	for n != 0 {
		if n&1 == 1 {
			ans = (a * ans) % mod
		}
		a = (a * a) % mod
		n = n >> 1
	}
	return ans
}
