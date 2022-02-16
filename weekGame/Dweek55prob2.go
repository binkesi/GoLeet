package weekgame

// https://leetcode-cn.com/contest/biweekly-contest-55/problems/remove-all-occurrences-of-a-substring/

func removeOccurrences(s string, part string) string {
	lp := len(part)
	bs, bp := []byte(s), []byte(part)
	var match func([]byte, int) bool
	match = func(nums []byte, i int) bool {
		for j := i; j < lp+i; j++ {
			if nums[j] != bp[j-i] {
				return false
			}
		}
		return true
	}
	for len(bs) >= lp {
		domatch := false
		for i := 0; i <= len(bs)-lp; i++ {
			if match(bs, i) {
				domatch = true
				bs = append(bs[0:i], bs[i+lp:]...)
				break
			}
		}
		if domatch == false {
			break
		}
	}
	return string(bs)
}
