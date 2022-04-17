package stringparse

import (
	"strings"
	"unicode"
)

// https://leetcode-cn.com/problems/most-common-word/

func mostCommonWord(paragraph string, banned []string) (ans string) {
	words := make([]string, 0)
	n := len(paragraph)
	for i, j := 0, 0; j < n; {
		for j < n && !unicode.IsLetter(rune(paragraph[i])) {
			i++
			j++
		}
		if j == n {
			break
		}
		for j < n && unicode.IsLetter(rune(paragraph[j])) {
			j++
		}
		words = append(words, strings.ToLower(paragraph[i:j]))
		i = j
	}
	maxCnt := 0
	wordMap := make(map[string]int)
	bannedMap := make(map[string]struct{})
	for _, word := range banned {
		bannedMap[word] = struct{}{}
	}
	for _, word := range words {
		if _, ok := bannedMap[word]; !ok {
			wordMap[word]++
			if wordMap[word] > maxCnt {
				maxCnt = wordMap[word]
				ans = word
			}
		}
	}
	return
}
