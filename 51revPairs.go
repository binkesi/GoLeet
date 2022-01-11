package goleet

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/

func reversePairs(nums []int) int {
	lenN := len(nums)
	if lenN == 0 {
		return 0
	}
	_, cnt := mergeSort(nums, 0, lenN-1)
	return cnt
}

func mergeSort(nums []int, start, end int) ([]int, int) {
	lena := end - start + 1
	if lena == 1 {
		return nums[start : end+1], 0
	} else {
		nums1, cnt1 := mergeSort(nums, start, start/2+end/2)
		nums2, cnt2 := mergeSort(nums, start/2+end/2+1, end)
		numsA, cnt := combine(nums1, nums2)
		return numsA, cnt1 + cnt2 + cnt
	}
}

func combine(nums1, nums2 []int) ([]int, int) {
	lena, lenb := len(nums1), len(nums2)
	var tmp []int = make([]int, 0)
	var cnt int = 0
	if lena == 1 && lenb == 1 {
		if nums1[0] <= nums2[0] {
			tmp = append(tmp, nums1[0], nums2[0])
			return tmp, cnt
		} else {
			tmp = append(tmp, nums2[0], nums1[0])
			cnt += 1
			return tmp, cnt
		}
	} else {
		ind1, ind2, contr := 0, 0, 0
		for ind1 != lena && ind2 != lenb {
			if nums1[ind1] <= nums2[ind2] {
				tmp = append(tmp, nums1[ind1])
				cnt += contr
				ind1 += 1
			} else {
				tmp = append(tmp, nums2[ind2])
				ind2 += 1
				contr += 1
			}
		}
		if ind1 != lena {
			tmp = append(tmp, nums1[ind1:]...)
			cnt += lenb * len(nums1[ind1:])
		} else {
			tmp = append(tmp, nums2[ind2:]...)
		}
	}
	return tmp, cnt
}
