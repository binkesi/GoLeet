package weekgame

import (
	"strconv"
	"strings"
	"unicode"
)

// https://leetcode-cn.com/contest/weekly-contest-263/problems/check-if-numbers-are-ascending-in-a-sentence/

func areNumbersAscending(s string) bool {
	strs := strings.Fields(s)
	ints := []int{}
	for _, v := range strs {
		if unicode.IsDigit(rune(v[0])) {
			num, _ := strconv.Atoi(v)
			ints = append(ints, num)
		}
	}
	lenI := len(ints)
	if lenI <= 1 {
		return true
	}
	for i := 1; i < lenI; i++ {
		if ints[i] <= ints[i-1] {
			return false
		}
	}
	return true
}
