package dfsbfs

// https://leetcode.cn/problems/different-ways-to-add-parentheses/solution/

func diffWaysToCompute(expression string) []int {
	cs := []byte(expression)
	n := len(expression)
	var dfs func(l, r int) []int
	dfs = func(l, r int) []int {
		res := []int{}
		for i := l; i <= r; i++ {
			if cs[i] >= '0' && cs[i] <= '9' {
				continue
			}
			l1, l2 := dfs(l, i-1), dfs(i+1, r)
			for _, v1 := range l1 {
				for _, v2 := range l2 {
					cur := 0
					if cs[i] == '+' {
						cur = v1 + v2
					} else if cs[i] == '-' {
						cur = v1 - v2
					} else {
						cur = v1 * v2
					}
					res = append(res, cur)
				}
			}
		}
		if len(res) == 0 {
			cur := 0
			for i := l; i <= r; i++ {
				cur = cur*10 + int(cs[i]-'0')
			}
			res = append(res, cur)
		}
		return res
	}
	return dfs(0, n-1)
}
