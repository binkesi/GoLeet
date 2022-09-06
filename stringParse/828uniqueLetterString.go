package stringparse

// https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/

func uniqueLetterString(s string) (ans int) {
	subStr := make(map[rune][]int)
	for i, b := range s {
		subStr[b] = append(subStr[b], i)
	}
	for _, arr := range subStr {
		arr = append(append([]int{-1}, arr...), len(s))
		for i := 1; i < len(arr)-1; i++ {
			ans += (arr[i] - arr[i-1]) * (arr[i+1] - arr[i])
		}
	}
	return
}
