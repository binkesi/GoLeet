package weekgame

// https://leetcode-cn.com/contest/weekly-contest-286/problems/find-the-difference-of-two-arrays/

func findDifference(nums1 []int, nums2 []int) (ans [][]int) {
	nMap1, nMap2 := make(map[int]struct{}), make(map[int]struct{})
	res1, res2 := make([]int, 0), make([]int, 0)
	for _, v := range nums1 {
		nMap1[v] = struct{}{}
	}
	for _, v := range nums2 {
		nMap2[v] = struct{}{}
	}
	for k, _ := range nMap1 {
		if _, ok := nMap2[k]; !ok {
			res1 = append(res1, k)
		}
	}
	for k, _ := range nMap2 {
		if _, ok := nMap1[k]; !ok {
			res2 = append(res2, k)
		}
	}
	ans = append(ans, res1, res2)
	return
}
