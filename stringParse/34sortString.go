// https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/
package stringparse

func searchRange(nums []int, target int) []int {
	length := len(nums)
	p1, p2 := 0, length-1
	if length == 0 || target < nums[p1] || target > nums[p2] {
		return []int{-1, -1}
	}
	in := (p1 + p2) / 2
	for p1 < p2 {
		if nums[in] == target {
			break
		}
		if nums[in] < target {
			p1 = in + 1
		}
		if nums[in] > target {
			p2 = in - 1
		}
		in = (p1 + p2) / 2
	}
	if nums[in] != target {
		return []int{-1, -1}
	}
	var start, end int
	for i := 0; i <= in; i++ {
		if nums[in-i] != target {
			start = in - i + 1
			break
		}
		if i == in {
			start = 0
		}
	}
	for j := 0; j <= length-in-1; j++ {
		if nums[in+j] != target {
			end = in + j - 1
			break
		}
		if j == length-in-1 {
			end = length - 1
		}
	}
	return []int{start, end}
}
