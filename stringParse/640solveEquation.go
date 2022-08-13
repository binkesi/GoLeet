package stringparse

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/solve-the-equation/

func solveEquation(equation string) string {
	subStr := strings.Split(equation, "=")
	left, right := subStr[0], subStr[1]
	l, r := len(left), len(right)
	label, num, hasx, hasn := 1, 0, 0, 0
	mods, nums := 0, 0
	for i := 0; i < l; i++ {
		if left[i] <= '9' && left[i] >= '0' {
			num = num*10 + int(left[i]-'0')
			hasn = 1
		}
		if left[i] == 'x' {
			hasx = 1
		}
		if left[i] == '-' || i == l-1 {
			if hasx == 0 {
				nums -= label * num
			} else {
				if hasn != 0 {
					mods += label * num
				} else {
					mods += label
				}
			}
			label, num, hasx, hasn = -1, 0, 0, 0
		} else if left[i] == '+' || i == l-1 {
			if hasx == 0 {
				nums -= label * num
			} else {
				if hasn != 0 {
					mods += label * num
				} else {
					mods += label
				}
			}
			label, num, hasx, hasn = 1, 0, 0, 0
		}
		if i == l-1 {
			label = 1
		}
	}
	for i := 0; i < r; i++ {
		if right[i] <= '9' && right[i] >= '0' {
			num = num*10 + int(right[i]-'0')
			hasn = 1
		}
		if right[i] == 'x' {
			hasx = 1
		}
		if right[i] == '-' || i == r-1 {
			if hasx == 0 {
				nums += label * num
			} else {
				if hasn != 0 {
					mods -= label * num
				} else {
					mods -= label
				}
			}
			label, num, hasx, hasn = -1, 0, 0, 0
		} else if right[i] == '+' || i == r-1 {
			if hasx == 0 {
				nums += label * num
			} else {
				if hasn != 0 {
					mods -= label * num
				} else {
					mods -= label
				}
			}
			label, num, hasx, hasn = 1, 0, 0, 0
		}
	}
	if mods == 0 && nums == 0 {
		return "Infinite solutions"
	}
	if mods == 0 && nums != 0 {
		return "No solution"
	}
	return "x=" + strconv.Itoa(nums/mods)
}
