package goleet

// https://leetcode-cn.com/problems/contains-duplicate-ii/

func containsNearbyDuplicate(nums []int, k int) bool {
	var numsMap map[int]int = make(map[int]int)
	lenN := len(nums)
	for i := 0; i < lenN; i++ {
		if _, ok := numsMap[nums[i]]; !ok {
			numsMap[nums[i]] = i
		} else if i-numsMap[nums[i]] <= k {
			return true
		} else {
			numsMap[nums[i]] = i
		}
	}
	return false
}
