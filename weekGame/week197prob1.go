package weekgame

// https://leetcode-cn.com/contest/weekly-contest-197/problems/number-of-good-pairs/

func numIdenticalPairs(nums []int) int {
	res := make(map[int]int)
	for _, v := range nums {
		if _, ok := res[v]; !ok {
			res[v] = 1
		} else {
			res[v] += 1
		}
	}
	ans := 0
	for i := range res {
		value := res[i]
		ans += (value - 1) * value / 2
	}
	return ans
}
