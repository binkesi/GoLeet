package goleet

// https://leetcode-cn.com/problems/find-peak-element/

func findPeakElement(nums []int) int {
	lenN := len(nums)
	if lenN == 1 {
		return 0
	}
	for ind := 0; ind < lenN; {
		if isPeak(nums, ind) {
			return ind
		}
		if nums[ind+1] > nums[ind] {
			ind = ind + 1
		} else {
			ind = ind + 2
		}
	}
	return 0
}

func isPeak(nums []int, ind int) bool {
	lenN := len(nums)
	if ind == 0 {
		return nums[ind] > nums[1]
	}
	if ind == lenN-1 {
		return nums[ind] > nums[lenN-2]
	}
	if nums[ind] > nums[ind-1] && nums[ind] > nums[ind+1] {
		return true
	} else {
		return false
	}
}
