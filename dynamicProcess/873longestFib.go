package dynamicprocess

func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	indMap := make(map[int]int)
	for i := range arr {
		indMap[arr[i]] = i
	}
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0 && j+2 > ans; j-- {
			diff := arr[i] - arr[j]
			if diff >= arr[j] {
				break
			}
			if ind, ok := indMap[diff]; ok {
				dp[i][j] = max(3, dp[j][ind]+1)
				if ans < dp[i][j] {
					ans = dp[i][j]
				}
			}
		}
	}
	return
}
