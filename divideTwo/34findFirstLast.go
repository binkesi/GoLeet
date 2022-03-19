package dividetwo

// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) (ans []int) {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] == target {
			i1, i2 := mid, mid
			for i1 > 0 && nums[i1-1] == target {
				i1--
			}
			for i2 < n-1 && nums[i2+1] == target {
				i2++
			}
			ans = []int{i1, i2}
			return
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	ans = []int{-1, -1}
	return
}
