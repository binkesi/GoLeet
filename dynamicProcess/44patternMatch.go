package dynamicprocess

// https://leetcode-cn.com/problems/wildcard-matching/

var match bool = false

func isMatch(s string, p string) bool {
	byteMatch(s, p, 0, 0)
	return match
}

func byteMatch(s string, p string, inds int, indp int) {
	if indp == len(p) {
		if inds == len(s) {
			match = true
		}
		return
	}
	if p[indp] == '*' {
		for i := 0; i <= len(s)-inds; i++ {
			byteMatch(s, p, inds+i, indp+1)
		}
	} else if p[indp] == '?' || p[indp] == s[inds] {
		byteMatch(s, p, inds+1, indp+1)
	}
}
