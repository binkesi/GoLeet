package goleet

// https://leetcode-cn.com/problems/calculate-money-in-leetcode-bank/

func totalMoney(n int) int {
	nTime, nLeft := n/7, n%7
	return 28*nTime + 7*(nTime-1)*nTime/2 + (1+2*nTime+nLeft)*nLeft/2
}
