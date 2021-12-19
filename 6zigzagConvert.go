// https://leetcode-cn.com/problems/zigzag-conversion/
package goleet

func convert(s string, numRows int) string {
	length := len(s)
	if numRows == 1 || numRows >= length {
		return s
	}
	var res []byte
	jmp := 2*numRows - 2
	for row := 0; row <= numRows-1; row++ {
		if row == 0 {
			res = append(res, s[0])
			for index := 1; index*jmp <= length-1; index++ {
				res = append(res, s[index*jmp])
			}
			continue
		}
		if row != numRows-1 {
			res = append(res, s[row])
			for index := 1; ; index++ {
				point := row + index*jmp
				if point-2*row <= length-1 {
					res = append(res, s[point-2*row])
				} else {
					break
				}
				if point <= length-1 {
					res = append(res, s[point])
				} else {
					break
				}
			}
		}
		if row == numRows-1 {
			res = append(res, s[numRows-1])
			for index := 1; numRows-1+index*jmp <= length-1; index++ {
				res = append(res, s[numRows-1+index*jmp])
			}
		}
	}
	return string(res)
}
