package arraymanipulate

// https://leetcode-cn.com/problems/sum-of-subarray-ranges/

func subArrayRanges(nums []int) (ans int64) {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		up, down := nums[i], nums[i]
		for j := i + 1; j < n; j++ {
			if nums[j] > up {
				up = nums[j]
			}
			if nums[j] < down {
				down = nums[j]
			}
			ans += int64(up - down)
		}
	}
	return
}
