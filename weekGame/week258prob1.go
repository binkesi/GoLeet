package weekgame

// https://leetcode-cn.com/contest/weekly-contest-258/problems/reverse-prefix-of-word/

func reversePrefix(word string, ch byte) string {
	lenW := len(word)
	r := []rune(word)
	for i := 0; i < lenW; i++ {
		if word[i] == ch {
			for j := 0; j <= i/2; j++ {
				r[j], r[i-j] = r[i-j], r[j]
			}
			break
		}
	}
	return string(r)
}
