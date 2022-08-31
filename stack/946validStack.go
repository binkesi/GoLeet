package stack

// https://leetcode.cn/problems/validate-stack-sequences/submissions/

func validateStackSequences(pushed []int, popped []int) bool {
	pushTop, popTop := 0, 0
	tmpStack := []int{}
	l := len(popped)
	for popTop < l {
		if len(tmpStack) > 0 && tmpStack[len(tmpStack)-1] == popped[popTop] {
			tmpStack = tmpStack[0 : len(tmpStack)-1]
			popTop++
		} else if pushTop >= l {
			return false
		} else {
			tmpStack = append(tmpStack, pushed[pushTop])
			pushTop++
		}
	}
	return true
}
