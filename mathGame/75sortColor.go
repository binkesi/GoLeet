package mathgame

//https://leetcode-cn.com/problems/sort-colors/

func sortColors(nums []int) {
	lenN := len(nums)
	start := 0
	for piviot := 0; piviot <= 1; piviot++ {
		i, j := start, start
		for j < lenN {
			if nums[j] <= piviot {
				nums[i], nums[j] = nums[j], nums[i]
				i += 1
				j += 1
			} else {
				j += 1
			}
		}
		start = i
	}
}
