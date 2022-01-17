package goleet

// https://leetcode-cn.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	lenN := len(nums)
	for i := 0; i < lenN; i++ {
		if nums[i] <= 0 {
			nums[i] = lenN + 1
		}
	}
	for j := 0; j < lenN; j++ {
		if abs(nums[j]) <= lenN {
			nums[abs(nums[j])-1] = -abs(nums[abs(nums[j])-1])
		}
	}
	for k := 0; k < lenN; k++ {
		if nums[k] > 0 {
			return k + 1
		}
	}
	return lenN + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
