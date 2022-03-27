package weekgame

// https://leetcode-cn.com/contest/weekly-contest-286/problems/minimum-deletions-to-make-array-beautiful/

func minDeletion(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	res := make([]int, 0)
	res = append(res, nums[0])
	for i, j := 1, 1; i < n; i++ {
		if j%2 == 1 {
			if nums[i] != res[j-1] {
				res = append(res, nums[i])
				j++
			} else {
				continue
			}
		} else {
			res = append(res, nums[i])
			j++
		}
	}
	m := len(res)
	if m%2 == 0 {
		return n - m
	} else {
		return n - m + 1
	}
}
