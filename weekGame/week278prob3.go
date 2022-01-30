package weekgame

// https://leetcode-cn.com/contest/weekly-contest-278/problems/find-substring-with-given-hash-value/

func subStrHash(s string, power int, modulo int, k int, hashValue int) string {
	lenS := len(s)
	var ans string
	if k == lenS {
		return s
	}
	hash, p, i := 0, 1, lenS-1
	for ; i >= lenS-k; i-- {
		hash = (hash*power + int(s[i]&31)) % modulo
		p = (p * power) % modulo
	}
	if hash == hashValue {
		ans = s[lenS-k:]
	}
	for ; i >= 0; i-- {
		hash = (hash*power + int(s[i]&31) + modulo - p*int(s[i+k]&31)%modulo) % modulo
		if hash == hashValue {
			ans = s[i : i+k]
		}
	}
	return ans
}
