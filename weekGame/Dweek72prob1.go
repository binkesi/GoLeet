package weekgame

// https://leetcode-cn.com/contest/biweekly-contest-72/problems/count-equal-and-divisible-pairs-in-an-array/

func countPairs(nums []int, k int) (ans int) {
	lenN := len(nums)
	for i := 0; i < lenN; i++ {
		for j := i + 1; j < lenN; j++ {
			if nums[j] == nums[i] && (i*j)%k == 0 {
				ans += 1
			}
		}
	}
	return
}
