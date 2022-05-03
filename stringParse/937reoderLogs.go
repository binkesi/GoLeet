package stringparse

import (
	"sort"
	"strings"
	"unicode"
)

func reorderLogFiles(logs []string) []string {
	sort.SliceStable(logs, func(i, j int) bool {
		a, b := strings.SplitN(logs[i], " ", 2), strings.SplitN(logs[j], " ", 2)
		if unicode.IsDigit(rune(a[1][0])) && unicode.IsDigit(rune(b[1][0])) {
			return false
		} else if !unicode.IsDigit(rune(a[1][0])) && !unicode.IsDigit(rune(b[1][0])) {
			return a[1] < b[1] || a[1] == b[1] && logs[i] < logs[j]
		}
		return !unicode.IsDigit(rune(a[1][0]))
	})
	return logs
}
