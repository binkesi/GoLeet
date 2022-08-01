package mathgame

// https://leetcode.cn/problems/valid-square/

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	a := (p2[0]-p1[0])*(p3[0]-p1[0]) + (p2[1]-p1[1])*(p3[1]-p1[1])
	b := (p2[0]-p1[0])*(p4[0]-p1[0]) + (p2[1]-p1[1])*(p4[1]-p1[1])
	c := (p3[0]-p1[0])*(p4[0]-p1[0]) + (p3[1]-p1[1])*(p4[1]-p1[1])
	if a == 0 {
		if b == 0 || c == 0 {
			return false
		}
		d := (p2[0]-p4[0])*(p3[0]-p4[0]) + (p2[1]-p4[1])*(p3[1]-p4[1])
		l1, l2 := (p2[0]-p1[0])*(p2[0]-p1[0])+(p2[1]-p1[1])*(p2[1]-p1[1]), (p3[0]-p1[0])*(p3[0]-p1[0])+(p3[1]-p1[1])*(p3[1]-p1[1])
		if d == 0 && l1 != 0 && l1 == l2 {
			return true
		}
		return false
	}
	if b == 0 {
		if a == 0 || c == 0 {
			return false
		}
		d := (p2[0]-p3[0])*(p4[0]-p3[0]) + (p2[1]-p3[1])*(p4[1]-p3[1])
		l1, l2 := (p2[0]-p1[0])*(p2[0]-p1[0])+(p2[1]-p1[1])*(p2[1]-p1[1]), (p4[0]-p1[0])*(p4[0]-p1[0])+(p4[1]-p1[1])*(p4[1]-p1[1])
		if d == 0 && l1 != 0 && l1 == l2 {
			return true
		}
		return false
	}
	if c == 0 {
		if a == 0 || b == 0 {
			return false
		}
		d := (p3[0]-p2[0])*(p4[0]-p2[0]) + (p3[1]-p2[1])*(p4[1]-p2[1])
		l1, l2 := (p3[0]-p1[0])*(p3[0]-p1[0])+(p3[1]-p1[1])*(p3[1]-p1[1]), (p4[0]-p1[0])*(p4[0]-p1[0])+(p4[1]-p1[1])*(p4[1]-p1[1])
		if d == 0 && l1 != 0 && l1 == l2 {
			return true
		}
		return false
	}
	return false
}
