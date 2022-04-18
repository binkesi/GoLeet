package dfsbfs

// https://leetcode-cn.com/problems/lexicographical-numbers/

func lexicalOrder(n int) (ans []int) {
	var dfs func(num, limit int)
	dfs = func(num, limit int) {
		if num > limit {
			return
		}
		ans = append(ans, num)
		for i := 0; i <= 9; i++ {
			tmp := num*10 + i
			dfs(tmp, limit)
		}
	}
	for i := 1; i <= 9; i++ {
		dfs(i, n)
	}
	return
}
