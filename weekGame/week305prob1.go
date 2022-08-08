package weekgame

// https://leetcode.cn/contest/weekly-contest-305/problems/number-of-arithmetic-triplets/

func arithmeticTriplets(nums []int, diff int) (ans int) {
	nDict := make(map[int]struct{})
	for _, v := range nums {
		nDict[v] = struct{}{}
	}
	for _, v := range nums {
		if _, ok := nDict[v+diff]; ok {
			if _, tok := nDict[v+diff+diff]; tok {
				ans++
			}
		}
	}
	return
}
