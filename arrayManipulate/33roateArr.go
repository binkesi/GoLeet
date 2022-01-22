// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
package arraymanipulate

func search(nums []int, target int) int {
	var in int
	length := len(nums)
	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	for i := 0; i < length; i++ {
		if target == nums[i] {
			in = i
			break
		}
		if i == length-1 {
			return -1
		}
	}
	return in
}
