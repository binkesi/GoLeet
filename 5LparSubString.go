// https://leetcode-cn.com/problems/longest-palindromic-substring/
package goleet

func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}
	var longest int = 1
	var longbyte string = s[:1]
	for i, v := range s {
		if i == length-1 {
			break
		}
		value := v
		for index, val := range s[i+1:] {
			if val == value && isPalindrome(s[i:index+i+2]) && index+2 > longest {
				longest = index + 2
				longbyte = s[i : index+i+2]
			}
		}
	}
	return longbyte
}

func isPalindrome(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}
	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}
	return true
}
