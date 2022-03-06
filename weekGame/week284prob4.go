package weekgame

// https://leetcode-cn.com/problems/replace-non-coprime-numbers-in-array/

func replaceNonCoprimes(nums []int) []int {
	n := len(nums)
	nStack := []int{nums[0]}
	for i := 1; i < n; i++ {
		nStack = append(nStack, nums[i])
		for len(nStack) != 1 {
			a, b := nStack[len(nStack)-1], nStack[len(nStack)-2]
			if ok, v := notCoprime(a, b); ok {
				m := a * b / v
				nStack[len(nStack)-2] = m
				nStack = nStack[0 : len(nStack)-1]
			} else {
				break
			}
		}
	}
	return nStack
}

func notCoprime(a, b int) (bool, int) {
	for a%b != 0 {
		a, b = b, a%b
	}
	if b > 1 {
		return true, b
	} else {
		return false, b
	}
}
