package weekgame

import "strconv"

// https://leetcode.cn/contest/weekly-contest-291/problems/k-divisible-elements-subarrays/

func countDistinct(nums []int, k int, p int) int {
	n := len(nums)
	nMap := make(map[string]struct{})
	for i := 0; i < n; i++ {
		cnt := 0
		str := ""
		for j := i; j < n; j++ {
			if nums[j]%p == 0 {
				cnt++
			}
			if cnt > k {
				break
			}
			str += strconv.Itoa(nums[j]) + "|"
			nMap[str] = struct{}{}
		}
	}
	return len(nMap)
}
