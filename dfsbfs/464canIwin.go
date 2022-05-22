package dfsbfs

// https://leetcode.cn/problems/can-i-win/submissions/

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	n := (1 << maxChoosableInteger) - 1
	fs := [1 << 20][2]int{}
	var dfs func(int, int, int) int
	if desiredTotal == 0 {
		return true
	}
	dfs = func(state, cur, k int) int {
		if state == n && cur < desiredTotal {
			return -1
		}
		if fs[state][k%2] != 0 {
			return fs[state][k%2]
		}
		var hope int
		if k%2 == 0 {
			hope = 1
		} else {
			hope = -1
		}
		for i := 0; i < maxChoosableInteger; i++ {
			if ((state >> i) & 1) == 1 {
				continue
			}
			if cur+i+1 >= desiredTotal {
				fs[state][k%2] = hope
				return fs[state][k%2]
			}
			if dfs(state|(1<<i), cur+i+1, k+1) == hope {
				fs[state][k%2] = hope
				return fs[state][k%2]
			}
		}
		fs[state][k%2] = -hope
		return fs[state][k%2]
	}
	res := dfs(0, 0, 0)
	if res == 1 {
		return true
	} else {
		return false
	}
}
