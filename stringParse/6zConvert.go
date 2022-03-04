package stringparse

// https://leetcode-cn.com/problems/zigzag-conversion/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	lenS := len(s)
	res := []byte{}
	for i := 0; i < numRows; i++ {
		if i == 0 {
			for j := 0; j < lenS; {
				res = append(res, s[j])
				j = j + (numRows-2)*2 + 2
			}
			continue
		}
		if i == numRows-1 {
			for j := i; j < lenS; {
				res = append(res, s[j])
				j = j + (numRows-2)*2 + 2
			}
			break
		}
		for j := 0; ; {
			in1, in2 := j-i, j+i
			if in1 >= 0 && in1 < lenS {
				res = append(res, s[in1])
			}
			if in2 >= 0 && in2 < lenS {
				res = append(res, s[in2])
			}
			if in1 >= lenS || in2 >= lenS {
				break
			}
			j = j + (numRows-2)*2 + 2
		}
	}
	return string(res)
}
