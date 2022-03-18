package backtracking

// https://leetcode-cn.com/problems/combination-sum/

func combinationSum(candidates []int, target int) (ans [][]int) {
	var pick func(can []int, cur []int, tar int, sum int)
	pick = func(can, cur []int, tar int, sum int) {
		if sum == tar {
			ans = append(ans, cur)
			return
		}
		n := len(cur)
		for _, v := range can {
			if (n == 0 || v >= cur[n-1]) && sum+v <= tar {
				next := make([]int, n)
				copy(next, cur)
				next = append(next, v)
				pick(can, next, tar, sum+v)
			}
		}
	}
	pick(candidates, []int{}, target, 0)
	return
}
