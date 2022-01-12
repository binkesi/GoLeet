package goleet

import "fmt"

// https://leetcode-cn.com/problems/search-a-2d-matrix-ii/

func searchMatrix(matrix [][]int, target int) bool {
	l, c := len(matrix), len(matrix[0])
	var start, end int
	if target < matrix[0][0] || target > matrix[l-1][c-1] {
		return false
	}
	for i, cnt := 0, 0; i < l; i++ {
		if matrix[i][c-1] >= target && cnt == 0 {
			start = i
			cnt += 1
		}
		if matrix[i][0] > target {
			end = i - 1
			break
		}
	}
	if end == 0 {
		end = l - 1
	}
	fmt.Println(start, end)
	for j := start; j <= end; j++ {
		if binSearch(matrix[j], target) {
			return true
		}
	}
	return false
}

func binSearch(nums []int, target int) bool {
	lenM := len(nums)
	left, right := 0, lenM-1
	mid := (left + right) / 2
	for left <= right {
		if nums[mid] == target {
			return true
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return false
}
