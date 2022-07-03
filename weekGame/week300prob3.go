package weekgame

// https://leetcode.cn/contest/weekly-contest-300/problems/number-of-people-aware-of-a-secret/

func peopleAwareOfSecret(n int, delay int, forget int) (ans int) {
	dp := make([]int, forget+1)
	dp[forget] = 1
	for i := 2; i <= n; i++ {
		tmp := make([]int, forget+1)
		for k := 0; k <= forget; k++ {
			if k < forget {
				tmp[k] = dp[k+1]
			} else {
				for l := 2; l <= forget-delay+1; l++ {
					tmp[k] += dp[l] % 1000000007
				}
			}
		}
		copy(dp, tmp)
	}
	for i := 1; i <= forget; i++ {
		ans += dp[i] % 1000000007
	}
	return ans % 1000000007
}
