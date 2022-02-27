package arraymanipulate

import "strconv"

// https://leetcode-cn.com/problems/optimal-division/

func optimalDivision(nums []int) (ans string) {
	lenN := len(nums)
	if lenN == 1 {
		ans = strconv.Itoa(nums[0])
		return
	}
	if lenN == 2 {
		ans = strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
		return
	}
	ans = strconv.Itoa(nums[0]) + "/(" + strconv.Itoa(nums[1])
	for i := 2; i < lenN; i++ {
		ans = ans + "/" + strconv.Itoa(nums[i])
	}
	ans = ans + ")"
	return
}
