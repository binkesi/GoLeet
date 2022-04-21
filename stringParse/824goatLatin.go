package stringparse

import "strings"

// https://leetcode-cn.com/problems/goat-latin/

func toGoatLatin(sentence string) (ans string) {
	substrs := strings.Split(sentence, " ")
	vowels := map[byte]struct{}{'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{},
		'o': struct{}{}, 'u': struct{}{}, 'A': struct{}{}, 'E': struct{}{}, 'I': struct{}{},
		'O': struct{}{}, 'U': struct{}{}}
	suffix := "a"
	for i, v := range substrs {
		if _, ok := vowels[v[0]]; ok {
			v = v + "ma"
		} else {
			v = v[1:] + string(v[0]) + "ma"
		}
		v = v + suffix
		suffix += "a"
		substrs[i] = v
	}
	var b strings.Builder
	n := len(substrs)
	for i, str := range substrs {
		b.WriteString(str)
		if i != n-1 {
			b.WriteString(" ")
		}
	}
	ans = b.String()
	return
}
