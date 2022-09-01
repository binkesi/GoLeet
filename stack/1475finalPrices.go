package stack

// https://leetcode.cn/problems/final-prices-with-a-special-discount-in-a-shop/

func finalPrices(prices []int) []int {
	myStack := []int{}
	ans := make([]int, len(prices))
next:
	for i := len(prices) - 1; i >= 0; i-- {
		for len(myStack) > 0 {
			if prices[i] >= myStack[len(myStack)-1] {
				ans[i] = prices[i] - myStack[len(myStack)-1]
				myStack = append(myStack, prices[i])
				continue next
			}
			myStack = myStack[0 : len(myStack)-1]
		}
		ans[i] = prices[i]
		myStack = append(myStack, prices[i])
	}
	return ans
}
