package slidingwindow

// https://leetcode.cn/problems/subarray-product-less-than-k/submissions/

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		res := 1
		for j := i; j < n; j++ {
			res *= nums[j]
			if res < k {
				ans++
			} else {
				break
			}
		}
	}
	return
}

func numSubarrayProductLessThanK2(nums []int, k int) (ans int) {
	i, prod := 0, 1
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += (j - i + 1)
	}
	return
}
