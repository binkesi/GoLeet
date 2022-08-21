package stringparse

import "strings"

// https://leetcode.cn/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/solution/

func isPrefixOfWord(sentence string, searchWord string) int {
	strs := strings.Split(sentence, " ")
	for i, s := range strs {
		if strings.HasPrefix(s, searchWord) {
			return i + 1
		}
	}
	return -1
}
