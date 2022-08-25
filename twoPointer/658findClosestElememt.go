package twopointer

import "sort"

func findClosestElements(arr []int, k int, x int) (ans []int) {
	l := len(arr)
	if x <= arr[0] {
		return arr[0:k]
	}
	if x >= arr[l-1] {
		return arr[l-k : l]
	}
	left, right := 0, l-1
	for left < right-1 {
		mid := left + (right-left)>>1
		if arr[mid] <= x {
			left = mid
		} else {
			right = mid
		}
	}
	for i := 0; i < k; i++ {
		if left < 0 {
			ans = append(ans, arr[right])
			right++
			continue
		}
		if right > l-1 {
			ans = append(ans, arr[left])
			left--
			continue
		}
		if x-arr[left] <= arr[right]-x {
			ans = append(ans, arr[left])
			left--
		} else {
			ans = append(ans, arr[right])
			right++
		}
	}
	sort.Ints(ans)
	return
}
