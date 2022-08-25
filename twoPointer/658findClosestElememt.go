package twopointer

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
			right++
			continue
		}
		if right > l-1 {
			left--
			continue
		}
		if x-arr[left] <= arr[right]-x {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}
