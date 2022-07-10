package weekgame

import "sort"

// https://leetcode.cn/contest/weekly-contest-301/problems/minimum-amount-of-time-to-fill-cups/

func fillCups(amount []int) (ans int) {
	sort.Ints(amount)
	if amount[0]+amount[1] >= amount[2] {
		ans += (amount[0]+amount[1]+amount[2])/2 + (amount[0]+amount[1]+amount[2])%2
	} else {
		ans += amount[0] + amount[1] + (amount[2] - amount[0] - amount[1])
	}
	return
}
