package dividetwo

// https://leetcode-cn.com/problems/single-element-in-a-sorted-array/

func singleNonDuplicate(nums []int) int {
	lenN := len(nums)
	if lenN == 1 {
		return nums[0]
	}
	l, r := 0, lenN-1
	for l < r {
		mid := l + (r-l)/2
		if mid%2 == 0 {
			mid = mid + 1
		}
		if nums[mid] == nums[mid-1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[l]
}
