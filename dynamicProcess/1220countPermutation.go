package dynamicprocess

// https://leetcode-cn.com/problems/count-vowels-permutation/

func countVowelPermutation(n int) int {
	var dp, tmp []int
	if n == 1 {
		return 5
	}
	dp = []int{1, 1, 1, 1, 1}
	tmp = make([]int, 5)
	for i := 2; i <= n; i++ {
		tmp[0] = dp[1] + dp[2] + dp[4]
		tmp[1] = dp[0] + dp[2]
		tmp[2] = dp[1] + dp[3]
		tmp[3] = dp[2]
		tmp[4] = dp[2] + dp[3]
		for k := 0; k < 5; k++ {
			dp[k] = tmp[k] % 1000000007
		}
	}
	var res int = 0
	for j := 0; j < 5; j++ {
		res += dp[j]
	}
	return int(res % 1000000007)
}
