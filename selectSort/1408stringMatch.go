package selectsort

import (
	"sort"
	"strings"
)

// https://leetcode.cn/problems/string-matching-in-an-array/submissions/

func stringMatching(words []string) (ans []string) {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	l := len(words)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if strings.Contains(words[j], words[i]) {
				ans = append(ans, words[i])
				break
			}
		}
	}
	return
}
