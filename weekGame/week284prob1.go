package weekgame

// https://leetcode-cn.com/contest/weekly-contest-284/problems/find-all-k-distant-indices-in-an-array/

func findKDistantIndices(nums []int, key int, k int) (ans []int) {
	ind := []int{}
	for i, v := range nums {
		if v == key {
			ind = append(ind, i)
		}
	}
	n := len(ind)
	if n == 0 {
		return
	}
	for i := range nums {
		for _, v := range ind {
			if absTwo(i, v) <= k {
				ans = append(ans, i)
				break
			}
		}
	}
	return
}

func absTwo(a, b int) int {
	if a <= b {
		return b - a
	} else {
		return a - b
	}
}
