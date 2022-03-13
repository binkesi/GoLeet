package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-284/problems/maximize-the-topmost-element-after-k-moves/

func maximumTop(nums []int, k int) (ans int) {
	if k == 0 {
		ans = nums[0]
		return
	}
	n := len(nums)
	if n == 1 {
		if k%2 == 0 {
			ans = nums[0]
			return
		} else {
			ans = -1
			return
		}
	}
	if k == 1 {
		ans = nums[1]
		return
	}
	tmp := make([]int, n)
	copy(tmp, nums)
	sort.Ints(tmp)
	if k > n {
		ans = tmp[n-1]
		return
	}
	nMap := make(map[int]struct{})
	for _, v := range nums[k-1:] {
		nMap[v] = struct{}{}
	}
	nLeft := make([]int, k-1)
	copy(nLeft, nums)
	sort.Ints(nLeft)
	ans = nLeft[k-2]
	// for j, cnt := 0, 0; j <= k-2; j++{
	// 	if _, ok := nMap[nLeft[j]]; !ok{
	// 		if nLeft[j] > ans{
	// 			ans = nLeft[j]
	// 			cnt += 1
	// 		}
	// 	}else if cnt >= 2{
	// 		if nLeft[j] > ans{
	// 			ans = nLeft[j]
	// 		}
	// 	}
	// }
	if k < n {
		if nums[k] > ans {
			ans = nums[k]
		}
	}
	return
}
