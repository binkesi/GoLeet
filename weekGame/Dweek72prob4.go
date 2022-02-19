package weekgame

// https://leetcode-cn.com/contest/biweekly-contest-72/problems/count-good-triplets-in-an-array/

func goodTriplets(nums1 []int, nums2 []int) (ans int64) {
	n := len(nums1)
	nMap1, nMap2 := make(map[int]int), make(map[int]int)
	for i := 0; i < n; i++ {
		nMap1[nums1[i]] = i
	}
	for j := 0; j < n; j++ {
		nMap2[nums2[j]] = j
	}
	for x1 := 0; x1 <= n-3; x1++ {
		n1 := nums1[x1]
		y1 := nMap2[n1]
		if y1 > n-3 {
			continue
		} else if y1 <= x1 {
			for x2 := x1 + 1; x2 <= n-2; x2++ {
				n2 := nums1[x2]
				y2 := nMap2[n2]
				if y2 > n-2 || y2 < y1 {
					continue
				} else if y2 <= x2 {
					for x3 := x2 + 1; x3 <= n-1; x3++ {
						n3 := nums1[x3]
						y3 := nMap2[n3]
						if y3 < y2 {
							continue
						} else {
							ans += 1
						}
					}
				} else {
					for y3 := y2 + 1; y3 <= n-1; y3++ {
						n3 := nums2[y3]
						x3 := nMap1[n3]
						if x3 < x2 {
							continue
						} else {
							ans += 1
						}
					}
				}
			}
		} else {
			for y2 := y1 + 1; y2 <= n-2; y2++ {
				n2 := nums2[y2]
				x2 := nMap1[n2]
				if x2 > n-2 || x2 < x1 {
					continue
				} else if x2 <= y2 {
					for y3 := y2 + 1; y3 <= n-1; y3++ {
						n3 := nums2[y3]
						x3 := nMap1[n3]
						if x3 < x2 {
							continue
						} else {
							ans += 1
						}
					}
				} else {
					for x3 := x2 + 1; x3 <= n-1; x3++ {
						n3 := nums1[x3]
						y3 := nMap2[n3]
						if y3 < y2 {
							continue
						} else {
							ans += 1
						}
					}
				}
			}
		}
	}
	return
}
