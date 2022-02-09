package weekgame

import "sort"

// https://leetcode-cn.com/contest/weekly-contest-279/problems/sort-even-and-odd-indices-independently/

func sortEvenOdd(nums []int) []int {
	lenN := len(nums)
	numOdd, numEven := make([]int, 0), make([]int, 0)
	for i, v := range nums {
		if i%2 == 1 {
			numOdd = append(numOdd, v)
		} else {
			numEven = append(numEven, v)
		}
	}
	sort.Ints(numEven)
	sort.Ints(numOdd)
	for i := 0; i < lenN; i++ {
		if i%2 == 1 {
			nums[i] = numOdd[lenN/2-1-i/2]
		} else {
			nums[i] = numEven[i/2]
		}
	}
	return nums
}
