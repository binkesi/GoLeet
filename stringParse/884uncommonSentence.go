package stringparse

import "strings"

// https://leetcode-cn.com/problems/uncommon-words-from-two-sentences/

func uncommonFromSentences(s1 string, s2 string) []string {
	words1, words2 := strings.Fields(s1), strings.Fields(s2)
	ans := []string{}
	wordMap1 := make(map[string]int)
	wordMap2 := make(map[string]int)
	for _, v := range words1 {
		if _, ok := wordMap1[v]; !ok {
			wordMap1[v] = 1
		} else {
			wordMap1[v] += 1
		}
	}
	for _, v := range words2 {
		if _, ok := wordMap2[v]; !ok {
			wordMap2[v] = 1
		} else {
			wordMap2[v] += 1
		}
	}
	for i := range wordMap1 {
		if wordMap1[i] == 1 {
			if _, ok := wordMap2[i]; !ok {
				ans = append(ans, i)
			}
		}
	}
	for i := range wordMap2 {
		if wordMap2[i] == 1 {
			if _, ok := wordMap1[i]; !ok {
				ans = append(ans, i)
			}
		}
	}
	return ans
}
