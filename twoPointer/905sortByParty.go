package twopointer

// https://leetcode.cn/problems/sort-array-by-parity/submissions/

func sortArrayByParity(nums []int) []int {
	n := len(nums)
	i, j := 0, 0
	for j < n {
		if nums[j]%2 == 1 {
			j++
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		}
	}
	return nums
}
