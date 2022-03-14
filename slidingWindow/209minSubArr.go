package slidingwindow

// https://leetcode-cn.com/problems/minimum-size-subarray-sum/

func minSubArrayLen(target int, nums []int) (ans int) {
	n := len(nums)
	sum, minL := 0, 0
	start, end := 0, 0
	for end < n {
		sum += nums[end]
		if sum >= target {
			minL = end + 1
			break
		}
		end++
	}
	if minL == 0 {
		ans = 0
		return
	}
	for end < n {
		curL := end - start + 1
		for sum >= target {
			if curL < minL {
				minL = curL
			}
			sum -= nums[start]
			curL--
			start++
		}
		for sum < target {
			end++
			if end < n {
				sum += nums[end]
				curL++
			} else {
				break
			}
		}
	}
	return minL
}
