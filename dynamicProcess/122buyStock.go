package dynamicprocess

// https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	maxp := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxp += (prices[i] - prices[i-1])
		}
	}
	return maxp
}
