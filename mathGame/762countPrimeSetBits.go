package mathgame

// https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/

func countPrimeSetBits(left int, right int) (ans int) {
	pMap := make(map[int]struct{})
	for i := left; i <= right; i++ {
		cnt := 0
		tmp := i
		for tmp != 0 {
			if tmp&1 == 1 {
				cnt++
			}
			tmp >>= 1
		}
		if _, ok := pMap[cnt]; ok {
			ans++
			continue
		} else {
			if isPrime(cnt) {
				ans++
				pMap[cnt] = struct{}{}
			}
		}
	}
	return
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
