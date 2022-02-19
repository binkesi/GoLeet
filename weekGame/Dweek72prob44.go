package weekgame

// https://leetcode-cn.com/contest/biweekly-contest-72/problems/count-good-triplets-in-an-array/

func goodTriplets2(nums1 []int, nums2 []int) (ans int64) {
	n := len(nums1)
	for i := 0; i < n; i++ {
		if nums1[i] == i {
			continue
		} else {
			ind := nums1[i]
			nums1[i], nums1[ind] = nums1[ind], nums1[i]
			nums2[i], nums2[ind] = nums2[ind], nums2[i]
		}
	}
	for i1 := 0; i1 <= n-3; i1++ {
		n1 := nums2[i1]
		for i2 := i1 + 1; i2 <= n-2; i2++ {
			n2 := nums2[i2]
			if n2 < n1 {
				continue
			} else {
				for i3 := i2 + 1; i3 <= n-1; i3++ {
					n3 := nums2[i3]
					if n3 > n2 {
						ans += 1
					}
				}
			}
		}
	}
	return
}
