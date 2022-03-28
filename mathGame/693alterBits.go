package mathgame

// https://leetcode-cn.com/problems/binary-number-with-alternating-bits/

func hasAlternatingBits(n int) bool {
	var next int
	if n&1 == 0 {
		next = 1
	} else {
		next = 0
	}
	n >>= 1
	for n > 0 {
		if next == 1 {
			if n&1 == 0 {
				return false
			}
			next = 0
		} else {
			if n&1 == 1 {
				return false
			}
			next = 1
		}
		n >>= 1
	}
	return true
}
