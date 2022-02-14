package weekgame

// https://leetcode-cn.com/contest/biweekly-contest-71/problems/minimum-cost-to-set-cooking-time/

func minCostSetTime(startAt int, moveCost int, pushCost int, targetSeconds int) (ans int) {
	m1, s1 := targetSeconds/60, targetSeconds%60
	res, res2 := []int{}, []int{}
	cost, cost2 := 0, 0
	if m1 <= 99 {
		if m1 > 0 {
			if m1/10 != 0 {
				res = append(res, m1/10, m1%10, s1/10, s1%10)
			} else {
				res = append(res, m1%10, s1/10, s1%10)
			}
		} else {
			if s1/10 != 0 {
				res = append(res, s1/10, s1%10)
			} else {
				res = append(res, s1%10)
			}
		}
		for i, beg := 0, startAt; i < len(res); i++ {
			if beg == res[i] {
				cost += pushCost
			} else {
				cost = cost + moveCost + pushCost
				beg = res[i]
			}
		}
	}
	if s1 <= 39 && m1 >= 1 {
		m2, s2 := m1-1, s1+60
		if m2 > 0 {
			if m2/10 != 0 {
				res2 = append(res2, m2/10, m2%10, s2/10, s2%10)
			} else {
				res2 = append(res2, m2%10, s2/10, s2%10)
			}
		} else {
			if s2/10 != 0 {
				res2 = append(res2, s2/10, s2%10)
			} else {
				res2 = append(res2, s2%10)
			}
		}
		for i, beg := 0, startAt; i < len(res2); i++ {
			if beg == res2[i] {
				cost2 += pushCost
			} else {
				cost2 = cost2 + moveCost + pushCost
				beg = res2[i]
			}
		}
	}
	if cost2 == 0 {
		ans = cost
		return
	}
	if cost == 0 {
		ans = cost2
		return
	}
	if cost <= cost2 {
		ans = cost
	} else {
		ans = cost2
	}
	return
}
