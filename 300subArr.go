package goleet

// https://leetcode-cn.com/problems/longest-increasing-subsequence/

func lengthOfLISar(nums []int) int {
	lenn := len(nums)
	var l []int = make([]int, 0)
	l = append(l, 0, nums[0])
	maxN := l[1]
	for i := 1; i < lenn; i++ {
		if nums[i] > maxN {
			l = append(l, nums[i])
			maxN = nums[i]
		} else {
			ind := binSearch(l, nums[i])
			l[ind] = nums[i]
			if ind == len(l)-1 {
				maxN = nums[i]
			}
		}
	}
	return len(l) - 1
}

func binSearch(l []int, n int) int {
	lenl := len(l)
	left, right := 1, lenl-1
	if l[left] > n {
		return left
	}
	mid := (left + right) / 2
	for left < right {
		if l[mid] < n {
			left = mid + 1
		} else if l[mid] > n {
			right = mid
		} else {
			return mid
		}
		mid = (left + right) / 2
	}
	return mid
}
