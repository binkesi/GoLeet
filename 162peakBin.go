package goleet

func findPeakElementBin(nums []int) int {
	lenN := len(nums)
	if lenN == 1 {
		return 0
	}
	for ind, piviot1, piviot2 := lenN/2, 0, lenN-1; ; {
		if isPeakBin(nums, ind) {
			return ind
		} else if ind < lenN-1 && nums[ind+1] > nums[ind] {
			piviot1 = ind + 1
		} else if ind > 0 && nums[ind-1] > nums[ind] {
			piviot2 = ind - 1
		}
		ind = (piviot1 + piviot2) / 2
	}
}

func isPeakBin(nums []int, ind int) bool {
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
