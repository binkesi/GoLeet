package arraymanipulate

// https://leetcode-cn.com/problems/maximum-difference-between-increasing-elements/

func maximumDifference(nums []int) (ans int) {
	lenN := len(nums)
	ans = -1
	for l, h := 0, 1; h < lenN; {
		if nums[h] <= nums[l] {
			l = h
			h += 1
		} else {
			if nums[h]-nums[l] > ans {
				ans = nums[h] - nums[l]
			}
			h += 1
		}
	}
	return
}
