package backtracking

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) (ans [][]int) {
	var backtrack func(nums, arr []int, cur, target int)
	combMap := make(map[])
	backtrack = func(nums, arr []int, cur, target int) {
		if cur == target{
			ans = append(ans, arr)
			return
		}
		for _, v := range nums{
			if cur + v <= target{
				backtrack
			}
		}
	}
}