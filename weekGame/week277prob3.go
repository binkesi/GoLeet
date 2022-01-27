package weekgame

// https://leetcode-cn.com/contest/weekly-contest-277/problems/find-all-lonely-numbers-in-the-array/

func findLonely(nums []int) []int {
	lenN := len(nums)
	res := []int{}
	numMap := make(map[int]int)
	for i := 0; i < lenN; i++ {
		if _, ok := numMap[nums[i]]; ok {
			numMap[nums[i]] += 1
		} else {
			numMap[nums[i]] = 0
		}
	}

	for k, v := range numMap {
		if v >= 1 {
			continue
		}
		if _, ok := numMap[k+1]; ok {
			continue
		} else if _, ok := numMap[k-1]; ok {
			continue
		}
		res = append(res, k)
	}
	return res
}
