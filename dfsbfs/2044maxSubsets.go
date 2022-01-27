package dfsbfs

func countMaxOrSubsets(nums []int) int {
	lenN := len(nums)
	res, target := 0, 0
	for _, v := range nums {
		target |= v
	}
	var dfs func(int, int)
	dfs = func(idx, cur int) {
		if cur == target {
			res += 1 << (lenN - idx)
			return
		}
		if idx == lenN {
			return
		}
		dfs(idx+1, cur|nums[idx])
		dfs(idx+1, cur)
	}
	dfs(0, 0)
	return res
}
