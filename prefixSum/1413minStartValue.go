package prefixsum

// https://leetcode.cn/problems/minimum-value-to-get-positive-step-by-step-sum/

func minStartValue(nums []int) int {
	res := nums[0]
	l := len(nums)
	for i := 1; i < l; i++ {
		nums[i] = nums[i-1] + nums[i]
		if nums[i] < res {
			res = nums[i]
		}
	}
	if res >= 0 {
		return 1
	}
	return 1 - res
}
