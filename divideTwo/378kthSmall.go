package dividetwo

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/

func kthSmallest2(matrix [][]int, k int) int {
	sortArr := combinek(matrix)
	return sortArr[k-1]
}

func combinek(matrix [][]int) []int {
	lenM := len(matrix)
	if lenM == 1 {
		return matrix[0]
	}
	if lenM == 2 {
		return mergeTwo(matrix[0], matrix[1])
	}
	return mergeTwo(combinek(matrix[0:lenM/2]), combinek(matrix[lenM/2:]))
}

func mergeTwo(nums1, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	var res []int = make([]int, 0, len1+len2)
	ind1, ind2 := 0, 0
	for ind1 < len1 && ind2 < len2 {
		if nums1[ind1] <= nums2[ind2] {
			res = append(res, nums1[ind1])
			ind1 += 1
		} else {
			res = append(res, nums2[ind2])
			ind2 += 1
		}
	}
	if ind1 == len1 {
		res = append(res, nums2[ind2:]...)
	} else {
		res = append(res, nums1[ind1:]...)
	}
	return res
}
