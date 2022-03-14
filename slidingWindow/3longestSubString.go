package slidingwindow

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	bMap := make(map[byte]int)
	bMap[s[0]] = 0
	start, end, length := 0, 1, 1
	maxLength := 1
	for end < n {
		if v, ok := bMap[s[end]]; !ok || v < start {
			bMap[s[end]] = end
			length += 1
			end++
		} else {
			start = bMap[s[end]] + 1
			bMap[s[end]] = end
			if length > maxLength {
				maxLength = length
			}
			length = end - start + 1
			end++
		}
	}
	if length > maxLength {
		maxLength = length
	}
	return maxLength
}
