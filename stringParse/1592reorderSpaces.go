package stringparse

import "strings"

// https://leetcode.cn/problems/rearrange-spaces-between-words/

func reorderSpaces(text string) string {
	cnt := strings.Count(text, " ")
	ans := strings.Builder{}
	l := len(text)
	words := []string{}
next:
	for i := 0; i < l; {
		if text[i] != ' ' {
			start := i
			for j := i + 1; j <= l; j++ {
				if j == l || text[j] == ' ' {
					end := j
					words = append(words, text[start:end])
					i = j
					continue next
				}
			}
		}
		i++
	}
	nWords := len(words)
	if nWords > 1 {
		nSpace := cnt / (nWords - 1)
		nLeft := cnt % (nWords - 1)
		for i, word := range words {
			if i != nWords-1 {
				ans.WriteString(word)
				ans.WriteString(strings.Repeat(" ", nSpace))
			} else {
				ans.WriteString(word)
				ans.WriteString(strings.Repeat(" ", nLeft))
			}
		}
	} else {
		ans.WriteString(words[0])
		ans.WriteString(strings.Repeat(" ", cnt))
	}
	return ans.String()
}
