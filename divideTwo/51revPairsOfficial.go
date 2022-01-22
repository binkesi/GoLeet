package dividetwo

// https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/

func reversePairsOff(nums []int) int {
	return mergeSortOff(nums, 0, len(nums)-1)
}

func mergeSortOff(nums []int, start, end int) int {
	if start >= end {
		return 0
	}
	mid := start/2 + end/2
	cnt := mergeSortOff(nums, start, mid) + mergeSortOff(nums, mid+1, end)
	var tmp []int = make([]int, 0)
	ind1, ind2 := start, mid+1
	for ind1 <= mid && ind2 <= end {
		if nums[ind1] <= nums[ind2] {
			tmp = append(tmp, nums[ind1])
			cnt += (ind2 - mid - 1)
			ind1 += 1
		} else {
			tmp = append(tmp, nums[ind2])
			ind2 += 1
		}
	}
	if ind1 > mid {
		tmp = append(tmp, nums[ind2:end+1]...)
	} else {
		tmp = append(tmp, nums[ind1:mid+1]...)
		cnt += (mid - ind1 + 1) * (end - mid)
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
	}
	return cnt
}
