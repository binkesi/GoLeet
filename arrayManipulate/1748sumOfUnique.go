package arraymanipulate

// https://leetcode-cn.com/problems/sum-of-unique-elements/

func sumOfUnique(nums []int) (ans int) {
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v] += 1
	}
	for k, v := range numMap {
		if v == 1 {
			ans += k
		}
	}
	return
}
