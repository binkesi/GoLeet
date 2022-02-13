package weekgame

import "math/bits"

// https://leetcode-cn.com/problems/maximum-and-sum-of-array/

func maximumANDSum(nums []int, numSlots int) (ans int) {
	lenN := len(nums)
	stats := make([]int, 1<<(numSlots*2))
	stats[0] = 0
	for i, fi := range stats {
		c := bits.OnesCount(uint(i))
		if c >= lenN {
			continue
		}
		for j := 0; j < numSlots*2; j++ {
			if i>>j&1 == 0 {
				s := i | 1<<j
				stats[s] = maxN(stats[s], fi+(j/2+1)&nums[c])
				ans = maxN(ans, stats[s])
			}
		}
	}
	return
}
