package arraymanipulate

import (
	"math"
	"sort"
)

// https://leetcode-cn.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/

func minimumDifference(nums []int, k int) (ans int) {
	sort.Ints(nums)
	lenN := len(nums)
	ans = math.MaxInt32
	for i := 0; i < lenN-k+1; i++ {
		if nums[i+k-1]-nums[i] < ans {
			ans = nums[i+k-1] - nums[i]
		}
	}
	return
}
