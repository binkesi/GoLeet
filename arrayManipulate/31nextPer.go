// https://leetcode-cn.com/problems/next-permutation/
package arraymanipulate

func nextPermutation(nums []int) {
	l := len(nums)
	if l == 0 || l == 1 {
		return
	}
	i := l - 2
	for i >= 0 {
		if nums[i] < nums[i+1] {
			break
		}
		i--
	}
	if i >= 0 {
		k := l - 1
		for ; k > i; k-- {
			if nums[k] > nums[i] {
				nums[k], nums[i] = nums[i], nums[k]
				reverse(nums[i+1:])
				break
			}
		}
	}
	if i < 0 {
		reverse(nums)
	}
}

func reverse(nums []int) {
	l := len(nums)
	for i := 0; i < l/2; i++ {
		nums[i], nums[l-1-i] = nums[l-1-i], nums[i]
	}
}
