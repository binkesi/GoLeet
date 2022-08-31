package arraymanipulate

// https://leetcode.cn/problems/shuffle-the-array/

func shuffle(nums []int, n int) []int {
	ans := make([]int, len(nums))
	for i, j, k := 0, n, 0; i < n; {
		ans[k] = nums[i]
		k++
		ans[k] = nums[j]
		k++
		i++
		j++
	}
	return ans
}
