package mathgame

// https://leetcode.cn/problems/prime-arrangements/submissions/

func numPrimeArrangements(n int) int {
	if n == 1 {
		return 1
	}
	var count int
	for i := 2; i <= n; i++ {
		if isPrime2(i) {
			count++
		}
	}
	left := n - count
	return stepMulti(count) * stepMulti(left) % 1000000007
}

func stepMulti(num int) (ans int) {
	if num == 1 {
		return 1
	}
	ans = 1
	for i := 2; i <= num; i++ {
		ans = ans * i % 1000000007
	}
	return
}

func isPrime2(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}
	var up int
	up = num / 2
	for i := 2; i <= up; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
