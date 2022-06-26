package arraymanipulate

// https://leetcode.cn/problems/duplicate-zeros/solution/

func duplicateZeros(arr []int) {
	n := len(arr)
	k := 0
	var ind int
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			k++
		}
		if i+k >= n-1 {
			ind = i
			break
		}
	}
	for j := ind; j >= 0; j-- {
		if arr[j] == 0 {
			if j+k > n-1 {
				arr[j+k-1] = 0
			} else {
				arr[j+k] = 0
				arr[j+k-1] = 0
			}
			k--
		} else {
			arr[j+k] = arr[j]
		}
	}
}
