package slidingwindow

// https://leetcode-cn.com/problems/rotate-function/submissions/

func maxRotateFunction(nums []int) (ans int) {
	n, sum := len(nums), nums[0]
	if n == 1 {
		return 0
	}
	for i := 1; i < n; i++ {
		ans += i * nums[i]
		sum += nums[i]
		nums[i] += nums[i-1]
	}
	start := ans
	for j := 1; j < n; j++ {
		tmp := start
		tmp += (n*nums[j-1] - j*sum)
		if tmp > ans {
			ans = tmp
		}
	}
	return
}
