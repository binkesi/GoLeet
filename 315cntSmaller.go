package goleet

// https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/

func countSmaller(nums []int) []int {
	lenN := len(nums)
	var res []int = make([]int, lenN)
	if lenN == 1 {
		return res
	}
	for i := lenN - 2; i >= 0; i-- {
		for j := i + 1; j <= lenN-1; j++ {
			if nums[j] < nums[i] {
				res[i] += 1
			}
		}
	}
	return res
}
