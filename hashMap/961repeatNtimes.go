package hashmap

// https://leetcode.cn/problems/n-repeated-element-in-size-2n-array/submissions/

func repeatedNTimes(nums []int) (ans int) {
	nMap := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := nMap[v]; !ok {
			nMap[v] = struct{}{}
		} else {
			ans = v
			break
		}
	}
	return
}
