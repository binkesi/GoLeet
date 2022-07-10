package weekgame

import "math"

func idealArrays(n int, maxValue int) (ans int) {
	for i := 1; i <= maxValue; i++ {
		pw := count(i)
		res := n * int(math.Pow(2.0, float64(pw))) % 1000000007
		ans = (ans + res) % 1000000007
	}
	return
}

func count(a int) int {
	cnt := 0
	for i := 2; i <= a/2; i++ {
		if a%i == 0 {
			cnt++
		}
	}
	return cnt
}
