package goleet

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	total := len1 + len2
	if total%2 == 1 {
		return float64(getKth(nums1, nums2, total/2+1))
	} else {
		return float64((getKth(nums1, nums2, total/2) + getKth(nums1, nums2, total/2+1))) / 2.0
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
		half := k / 2
		newind1 := min(index1+half, len(nums1)) - 1
		newind2 := min(index2+half, len(nums2)) - 1
		if nums1[newind1] <= nums2[newind2] {
			k -= (newind1 - index1 + 1)
			index1 = newind1 + 1
		} else {
			k -= (newind2 - index2 + 1)
			index2 = newind2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
