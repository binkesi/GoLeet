package goleet

// https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self/

func countSmaller(nums []int) []int {
	var revCnt []int = make([]int, len(nums))
	var index []int = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		index[i] = i
	}
	mergeSortAll(nums, index, revCnt, 0, len(nums)-1)
	return revCnt
}

func mergeSortAll(nums, index, revCnt []int, start, end int) {
	if start >= end {
		return
	}
	var tmp []int = make([]int, 0)
	var tmpInd []int = make([]int, 0)
	mid := start/2 + end/2
	ind1, ind2 := start, mid+1
	mergeSortAll(nums, index, revCnt, start, mid)
	mergeSortAll(nums, index, revCnt, mid+1, end)
	for ind1 <= mid && ind2 <= end {
		if nums[ind1] <= nums[ind2] {
			tmp = append(tmp, nums[ind1])
			tmpInd = append(tmpInd, index[ind1])
			revCnt[index[ind1]] += (ind2 - (mid + 1))
			ind1 += 1
		} else {
			tmp = append(tmp, nums[ind2])
			tmpInd = append(tmpInd, index[ind2])
			ind2 += 1
		}
	}
	if ind1 > mid {
		tmp = append(tmp, nums[ind2:end+1]...)
		for k := ind2; k <= end; k++ {
			tmpInd = append(tmpInd, index[k])
		}
	} else {
		tmp = append(tmp, nums[ind1:mid+1]...)
		for k := ind1; k <= mid; k++ {
			tmpInd = append(tmpInd, index[k])
		}
		for i := ind1; i <= mid; i++ {
			revCnt[index[i]] += (end - mid)
		}
	}
	for i := start; i <= end; i++ {
		nums[i] = tmp[i-start]
		index[i] = tmpInd[i-start]
	}
}
