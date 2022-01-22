package stringparse

func myAtoi(s string) int {
	const INT_MAX = int(^uint32(0) >> 1)
	var mystring []int
	var mynum = 0
	var cha string
	for i, v := range s {
		if v == ' ' {
			if i == 0 || s[i-1] == ' ' {
				continue
			} else {
				break
			}
		}
		if v == '+' || v == '-' {
			if i == 0 || s[i-1] == ' ' {
				cha = string(v)
			} else {
				break
			}
			continue
		}
		d := v - '0'
		if d >= 0 && d <= 9 {
			mystring = append(mystring, int(d))
		} else {
			break
		}
	}
	for i := 0; i < len(mystring); i++ {
		if mynum*10 > INT_MAX {
			mynum = INT_MAX + 2
			break
		}
		mynum = mynum*10 + mystring[i]
	}
	if cha == "" || cha == "+" {
		if mynum > INT_MAX {
			return INT_MAX
		} else {
			return mynum
		}
	} else {
		if mynum > INT_MAX+1 {
			return INT_MAX*(-1) - 1
		} else {
			return mynum * (-1)
		}
	}
}
