package weekgame

// https://leetcode-cn.com/contest/biweekly-contest-71/problems/partition-array-according-to-given-pivot/

func pivotArray(nums []int, pivot int) (ans []int) {
	lenN := len(nums)
	a, b, c := []int{}, []int{}, []int{}
	for i := 0; i < lenN; i++ {
		if nums[i] < pivot {
			a = append(a, nums[i])
		} else if nums[i] == pivot {
			b = append(b, nums[i])
		} else {
			c = append(c, nums[i])
		}
	}
	ans = append(a, append(b, c...)...)
	return
}
