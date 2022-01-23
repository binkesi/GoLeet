package graph

import "fmt"

// https://leetcode-cn.com/problems/word-ladder/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	convertMap := make(map[int][]int)
	lenW := len(wordList)
	for i, j := 0, 0; i < lenW; i++ {
		if wordList[i] == endWord {
			j = 1
		} else if ableConvert(wordList[i], endWord) {
			convertMap[i] = append(convertMap[i], lenW)
		}
		if i == lenW-1 && j == 0 {
			return 0
		}
	}
	for i := 0; i < lenW; i++ {
		if ableConvert(beginWord, wordList[i]) {
			convertMap[-1] = append(convertMap[-1], i)
		}
		if ableConvert(beginWord, endWord) {
			convertMap[-1] = append(convertMap[-1], lenW)
		}
	}
	if len(convertMap[-1]) == 0 {
		return 0
	}
	for i := 0; i < lenW; i++ {
		for j := i + 1; j < lenW; j++ {
			if ableConvert(wordList[i], wordList[j]) {
				convertMap[i] = append(convertMap[i], j)
				convertMap[j] = append(convertMap[j], i)
			}
		}
	}
	for i := 0; i < lenW; i++ {
		if wordList[i] == endWord {
			delete(convertMap, i)
		}
	}
	fmt.Println(convertMap)
	type pair struct{ idx, step int }
	queue := []pair{{idx: -1, step: 1}}
	visited := map[int]bool{-1: true}
	for len(queue) != 0 {
		q := queue[0]
		queue = queue[1:]
		idx, step := q.idx, q.step
		if idx == lenW {
			return step
		}
		for _, v := range convertMap[idx] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, pair{idx: v, step: step + 1})
			}
		}
	}
	return 0
}

func ableConvert(a, b string) bool {
	lenA := len(a)
	var dif int = 0
	for i := 0; i < lenA; i++ {
		if a[i] != b[i] {
			dif += 1
		}
	}
	return dif == 1
}
