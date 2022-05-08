package arraymanipulate

// https://leetcode.cn/problems/find-all-duplicates-in-an-array/submissions/

func findDuplicates(nums []int) (ans []int) {
	n := len(nums)
next:
	for i := 0; i < n; i++ {
		if nums[i] == i+1 {
			continue
		}
		tmp := nums[i]
		ind := tmp - 1
		nums[i] = 0
		for tmp != nums[ind] {
			tmp = nums[ind]
			nums[ind] = ind + 1
			if tmp == 0 {
				continue next
			}
			ind = tmp - 1
		}
		ans = append(ans, tmp)
	}
	return
}
