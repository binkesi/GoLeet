package stringparse

import (
	"strings"
	"unicode"
)

// https://leetcode-cn.com/problems/number-of-valid-words-in-a-sentence/

func countValidWords(sentence string) int {
	lenS := len(sentence)
	words := []string{}
	for i := 0; i < lenS; {
		start, end := 0, 0
		if sentence[i] == ' ' {
			i += 1
			continue
		} else {
			start = i
			for i < lenS && sentence[i] != ' ' {
				i += 1
			}
			end = i
			words = append(words, sentence[start:end])
		}
	}
	res := 0
	for _, v := range words {
		if isValid(v) {
			res += 1
		}
	}
	return res
}

func isValid(word string) bool {
	hasHyphen := false
	for i, s := range word {
		if unicode.IsDigit(s) || strings.ContainsRune("!.,", s) && i < len(word)-1 {
			return false
		}
		if s == '-' {
			if hasHyphen || i == 0 || i == len(word)-1 || !unicode.IsLower(rune(word[i-1])) || !unicode.IsLower(rune(word[i+1])) {
				return false
			} else {
				hasHyphen = true
			}
		}
	}
	return true
}
