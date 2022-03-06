package weekgame

// https://leetcode-cn.com/contest/weekly-contest-283/problems/append-k-integers-with-minimal-sum/

func minimalKSum(nums []int, k int) (ans int64) {
	n := len(nums)
	ans = int64(k * (1 + k) / 2)
	addMap := make(map[int]struct{})
	nMap := make(map[int]struct{})
	for _, v := range nums {
		nMap[v] = struct{}{}
	}
	for i, j := 0, 1; i < n; i++ {
		if _, ok := addMap[nums[i]]; nums[i] <= k && !ok {
			ans = ans - int64(nums[i])
			addMap[nums[i]] = struct{}{}
			for {
				if _, ok := nMap[k+j]; ok {
					j += 1
				} else {
					break
				}
			}
			ans = ans + int64(k+j)
			j++
		}
	}
	return
}
