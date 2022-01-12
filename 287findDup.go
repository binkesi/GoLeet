package goleet

// https://leetcode-cn.com/problems/find-the-duplicate-number/

func findDuplicate(nums []int) int {
	lenN := len(nums)
	left, right := 1, lenN-1
	mid := (right + left) >> 1
	ans := 1
	for left < right {
		sumN := 0
		for i := 0; i < lenN; i++ {
			if nums[i] <= mid {
				sumN += 1
			}
		}
		if sumN <= mid {
			left = mid + 1
		} else {
			right = mid
		}
		mid = (right + left) >> 1
		ans = mid
	}
	return ans
}
