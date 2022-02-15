package weekgame

// https://leetcode-cn.com/contest/weekly-contest-262/problems/two-out-of-three/

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) (ans []int) {
	numMap := make(map[int]map[int]struct{})
	l1, l2, l3 := len(nums1), len(nums2), len(nums3)
	for i := 0; i < l1; i++ {
		if _, ok := numMap[nums1[i]]; !ok {
			numMap[nums1[i]] = make(map[int]struct{})
			numMap[nums1[i]][1] = struct{}{}
		}
	}
	for j := 0; j < l2; j++ {
		if _, ok := numMap[nums2[j]]; ok {
			numMap[nums2[j]][2] = struct{}{}
		} else {
			numMap[nums2[j]] = make(map[int]struct{})
			numMap[nums2[j]][2] = struct{}{}
		}
	}
	for k := 0; k < l3; k++ {
		if _, ok := numMap[nums3[k]]; ok {
			numMap[nums3[k]][3] = struct{}{}
		} else {
			numMap[nums3[k]] = make(map[int]struct{})
			numMap[nums3[k]][3] = struct{}{}
		}
	}
	for k, v := range numMap {
		if len(v) >= 2 {
			ans = append(ans, k)
		}
	}
	return
}
