package mathgame

// https://leetcode.cn/problems/consecutive-numbers-sum/submissions/

func consecutiveNumbersSum(n int) (ans int) {
	n *= 2
	for k := 1; k*k < n; k++ {
		if n%k != 0 {
			continue
		}
		if (n/k+1-k)%2 == 0 {
			ans++
		}
	}
	return
}
