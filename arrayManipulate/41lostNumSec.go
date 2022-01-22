package arraymanipulate

// https://leetcode-cn.com/problems/first-missing-positive/

func firstMissingPositive(nums []int) int {
	lenN := len(nums)
	for i := 0; i < lenN; i++ {
		for nums[i] > 0 && nums[i] <= lenN && nums[i] != i+1 {
			tmp := nums[nums[i]-1]
			if tmp == nums[i] {
				break
			}
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		}
	}
	for j := 0; j < lenN; j++ {
		if nums[j] != j+1 {
			return j + 1
		}
	}
	return lenN + 1
}
