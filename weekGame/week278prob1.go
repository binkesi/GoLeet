package weekgame

// https://leetcode-cn.com/contest/weekly-contest-278/problems/keep-multiplying-found-values-by-two/

func findFinalValue(nums []int, original int) int {
	numMap := make(map[int]struct{})
	for _, v := range nums {
		numMap[v] = struct{}{}
	}
	for {
		if _, ok := numMap[original]; ok {
			original *= 2
		} else {
			break
		}
	}
	return original
}
