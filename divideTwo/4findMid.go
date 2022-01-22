package dividetwo

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lena, lenb := len(nums1), len(nums2)
	total := lena + lenb
	if total%2 == 1 {
		k := total/2 + 1
		return float64(getKth(nums1, nums2, k))
	} else {
		k1, k2 := total/2, total/2+1
		return float64(getKth(nums1, nums2, k1)+getKth(nums1, nums2, k2)) / 2
	}
}

func getKth(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		newInd1 := min(index1+k/2, len(nums1)) - 1
		newInd2 := min(index2+k/2, len(nums2)) - 1
		if nums1[newInd1] < nums2[newInd2] {
			k -= (newInd1 - index1 + 1)
			index1 = newInd1 + 1
		} else {
			k -= (newInd2 - index2 + 1)
			index2 = newInd2 + 1
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
