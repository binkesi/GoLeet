package slidingwindow

// https://leetcode-cn.com/problems/maximum-average-subarray-i/

func findMaxAverage(nums []int, k int) (ans float64) {
	n := len(nums)
	sum, start, end := 0, 0, k-1
	for i := start; i <= end; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for end < n-1 {
		end++
		sum += nums[end]
		sum -= nums[start]
		start++
		if sum > maxSum {
			maxSum = sum
		}
	}
	ans = float64(maxSum) / float64(k)
	return
}
