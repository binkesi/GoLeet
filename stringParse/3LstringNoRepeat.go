// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
package stringparse

func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}
	strmap := make(map[string]int)
	var longest int = 0
	var thislen int = 0
	for k, v := range s {
		value := string(v)
		if val, ok := strmap[value]; ok {
			for m, n := range strmap {
				if n <= val {
					delete(strmap, m)
				}
			}
			if thislen > longest {
				longest = thislen
			}
			strmap[value] = k
			thislen = k - val
			continue
		}
		thislen += 1
		strmap[value] = k
	}
	if thislen > longest {
		longest = thislen
	}
	return longest
}
