package priorityqueue

import (
	"math/rand"
	"time"
)

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

func findKthLargest2(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return qucikSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func qucikSelect(nums []int, l, r, ind int) int {
	q := randPartion(nums, l, r)
	if q == ind {
		return nums[q]
	} else if q < ind {
		return qucikSelect(nums, q+1, r, ind)
	}
	return qucikSelect(nums, l, q-1, ind)
}

func randPartion(nums []int, l, r int) int {
	i := rand.Int()%(r-l+1) + l
	nums[i], nums[r] = nums[r], nums[i]
	return partion(nums, l, r)
}

func partion(nums []int, l, r int) int {
	num := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] < num {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
