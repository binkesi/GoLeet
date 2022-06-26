package weekgame

// https://leetcode.cn/contest/weekly-contest-299/problems/maximum-score-of-spliced-array/

func maximumsSplicedArray(nums1 []int, nums2 []int) int {
	sum1, sum2 := 0, 0
	n := len(nums1)
	dp1, dp2 := make([]int, n), make([]int, n)
	max1, max2 := 0, 0
	for i := 0; i < n; i++ {
		sum1 += nums1[i]
		sum2 += nums2[i]
	}
	if nums1[0] > nums2[0] {
		dp1[0] = nums1[0] - nums2[0]
		max1 = dp1[0]
	} else {
		dp2[0] = nums2[0] - nums1[0]
		max2 = dp2[0]
	}
	for i := 1; i < n; i++ {
		if nums1[i]-nums2[i]+dp1[i-1] > 0 {
			dp1[i] = nums1[i] - nums2[i] + dp1[i-1]
		} else {
			dp1[i] = 0
		}
		if dp1[i] > max1 {
			max1 = dp1[i]
		}
	}
	for i := 1; i < n; i++ {
		if nums2[i]-nums1[i]+dp2[i-1] > 0 {
			dp2[i] = nums2[i] - nums1[i] + dp2[i-1]
		} else {
			dp2[i] = 0
		}
		if dp2[i] > max2 {
			max2 = dp2[i]
		}
	}
	return max(sum2+max1, sum1+max2)
}
