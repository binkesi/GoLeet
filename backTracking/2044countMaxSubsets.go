package backtracking

// https://leetcode-cn.com/problems/count-number-of-maximum-bitwise-or-subsets/

func countMaxOrSubsets(nums []int) (ans int) {
	n := len(nums)
	up := (1 << n) - 1
	res := 0
	for _, v := range nums {
		res |= v
	}
	for i := 1; i <= up; i++ {
		tmp := i
		rsl := 0
		for k := 0; tmp != 0; k++ {
			if tmp&1 == 1 {
				rsl |= nums[n-k-1]
			}
			if rsl == res {
				ans += 1
				break
			}
			tmp = tmp >> 1
		}
	}
	return
}

func countMaxOrSubsets2(nums []int) (ans int) {
	n := len(nums)
	res := 0
	for _, v := range nums {
		res |= v
	}
	var dfs func(int, int)
	dfs = func(ind, rsl int) {
		if rsl == res {
			ans += 1 << (n - ind)
			return
		}
		if ind == n {
			return
		}
		dfs(ind+1, rsl|nums[ind])
		dfs(ind+1, rsl)
	}
	dfs(0, 0)
	return
}
