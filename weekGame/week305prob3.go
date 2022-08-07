package weekgame

// https://leetcode.cn/contest/weekly-contest-305/problems/check-if-there-is-a-valid-partition-for-the-array/

func validPartition(nums []int) bool {
	l := len(nums)
	f := make([]bool, l+1)
	f[0] = true
	for i, num := range nums {
		if i > 0 && f[i-1] && num == nums[i-1] || i > 1 && f[i-2] && ((num == nums[i-1] && num == nums[i-2]) || (num-1 == nums[i-1] && num-2 == nums[i-2])) {
			f[i+1] = true
		}
	}
	return f[l]
}
