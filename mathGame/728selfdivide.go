package mathgame

// https://leetcode-cn.com/problems/self-dividing-numbers/

func selfDividingNumbers(left int, right int) (ans []int) {
	var selfDivide func(int) bool
	selfDivide = func(i int) bool {
		num := i
		for num > 0 {
			tmp := num % 10
			if tmp == 0 || i%tmp != 0 {
				return false
			}
			num /= 10
		}
		return true
	}
	for i := left; i <= right; i++ {
		if selfDivide(i) {
			ans = append(ans, i)
		}
	}
	return
}
