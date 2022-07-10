package weekgame

// https://leetcode.cn/contest/weekly-contest-301/problems/move-pieces-to-obtain-a-string/

func canChange(start string, target string) bool {
	n := len(start)
	for p1, p2 := 0, 0; p1 < n && p2 < n; {
		for p1 < n && start[p1] == '_' {
			p1++
		}
		for p2 < n && target[p2] == '_' {
			p2++
		}
		if p1 == n && p2 == n {
			return true
		}
		if (p1 == n && p2 != n) || (p1 != n && p2 == n) {
			return false
		}
		if start[p1] != target[p2] {
			return false
		}
		if start[p1] == 'L' {
			p1++
			p2++
			continue
		}
		if p1 > p2 {
			return false
		}
		p1++
		p2++
	}

	for p1, p2 := n-1, n-1; p1 >= 0 && p2 >= 0; {
		for p1 >= 0 && start[p1] == '_' {
			p1--
		}
		for p2 >= 0 && target[p2] == '_' {
			p2--
		}
		if p1 < 0 && p2 < 0 {
			return true
		}
		if (p1 < 0 && p2 >= 0) || (p1 >= 0 && p2 < 0) {
			return false
		}
		if start[p1] != target[p2] {
			return false
		}
		if start[p1] == 'R' {
			p1--
			p2--
			continue
		}
		if p1 < p2 {
			return false
		}
		p1--
		p2--
	}
	return true
}
