package dividetwo

func search(nums []int, target int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)>>1
		nm := nums[mid]
		if nm == target {
			return mid
		}
		if nm < nums[r] {
			if nm > nums[l] {
				if nm > target {
					r = mid - 1
				} else {
					l = mid + 1
				}
			} else if nums[r] < target || nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[l] > target || nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
