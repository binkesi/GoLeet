package arraymanipulate

import (
	"math/rand"
	"time"
)

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/submissions/

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	n := len(nums)
	return quickSelect(nums, 0, n-1, n-k)
}

func quickSelect(nums []int, l, r, ind int) int {
	q := randPartion(nums, l, r)
	if q == ind {
		return nums[q]
	} else if q < ind {
		return quickSelect(nums, q+1, r, ind)
	}
	return quickSelect(nums, l, q-1, ind)
}

func randPartion(nums []int, l, r int) int {
	ind := rand.Int()%(r-l+1) + l
	nums[r], nums[ind] = nums[ind], nums[r]
	return partion(nums, l, r)
}

func partion(nums []int, l, r int) int {
	num := nums[r]
	ind := l
	for j := l; j <= r; j++ {
		if nums[j] < num {
			nums[ind], nums[j] = nums[j], nums[ind]
			ind++
		}
	}
	nums[ind], nums[r] = nums[r], nums[ind]
	return ind
}
