package weekgame

// https://leetcode-cn.com/contest/weekly-contest-277/problems/rearrange-array-elements-by-sign/

func rearrangeArray(nums []int) []int {
	lenN := len(nums)
	arrpos, arrneg := make([]int, 0, lenN/2), make([]int, 0, lenN/2)
	for i := 0; i < lenN; i++ {
		if nums[i] > 0 {
			arrpos = append(arrpos, nums[i])
		} else {
			arrneg = append(arrneg, nums[i])
		}
	}
	var res []int = make([]int, 0, lenN)
	for k := 0; k < lenN/2; k++ {
		res = append(res, arrpos[k], arrneg[k])
	}
	return res
}
