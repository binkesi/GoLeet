package twopointer

import "math"

// https://leetcode.cn/problems/next-greater-element-iii/submissions/

func nextGreaterElement(n int) (ans int) {
	const up = math.MaxInt32
	if n <= 10 {
		return -1
	}
	nums := []int{}
	for n > 0 {
		nums = append(nums, n%10)
		n /= 10
	}
	m := len(nums)
	for i, j := 0, 1; j < m; {
		if nums[i] > nums[j] {
			k := i
			for k >= 0 && nums[k] > nums[j] {
				k--
			}
			nums[k+1], nums[j] = nums[j], nums[k+1]
			for l, r := 0, i; l < r; {
				nums[l], nums[r] = nums[r], nums[l]
				l++
				r--
			}
			p := 1
			for k := 0; k < m; k++ {
				ans += p * nums[k]
				p *= 10
			}
			if ans > up {
				return -1
			} else {
				return
			}
		}
		i++
		j++
	}
	return -1
}
