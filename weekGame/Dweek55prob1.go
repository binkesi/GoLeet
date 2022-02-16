package weekgame

// https://leetcode-cn.com/contest/biweekly-contest-55/problems/remove-one-element-to-make-the-array-strictly-increasing/

func canBeIncreasing(nums []int) bool {
	lenN := len(nums)
	var increase func([]int) bool
	increase = func(arr []int) bool {
		l := len(arr)
		for i := 1; i < l; i++ {
			if arr[i] <= arr[i-1] {
				return false
			}
		}
		return true
	}
	for j := 1; j < lenN; j++ {
		if nums[j] <= nums[j-1] {
			tmp1, tmp2 := make([]int, lenN), make([]int, lenN)
			copy(tmp1, nums)
			copy(tmp2, nums)
			return increase(append(tmp1[0:j], tmp1[j+1:]...)) || increase(append(tmp2[0:j-1], tmp2[j:]...))
		}
	}
	return true
}
