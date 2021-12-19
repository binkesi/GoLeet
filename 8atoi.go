// https://leetcode-cn.com/problems/string-to-integer-atoi/
package goleet

func myAtoi(s string) int {
	var b []byte
	var nums map[string]struct{} = make(map[string]struct{})
	nums = map[string]struct{}{
		"0": struct{}{},
		"1": struct{}{},
		"2": struct{}{},
		"3": struct{}{},
		"4": struct{}{},
		"5": struct{}{},
		"6": struct{}{},
		"7": struct{}{},
		"8": struct{}{},
		"9": struct{}{},
	}
	length := len(s)
	var res string
	for i := 0; i < length; i++ {
		if string(s[i]) == " " {
			continue
		}
		j := i
		if _, ok := nums[string(s[i])]; !ok {
			break
		}
	}
}
