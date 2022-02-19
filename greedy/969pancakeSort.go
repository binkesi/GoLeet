package greedy

// https://leetcode-cn.com/problems/pancake-sorting/

func pancakeSort(arr []int) (ans []int) {
	for n := len(arr); n > 1; n-- {
		ind := 0
		for i, v := range arr[:n] {
			if v > arr[ind] {
				ind = i
			}
		}
		if ind == n-1 {
			continue
		}
		for i, m := 0, ind; i < (m+1)/2; i++ {
			arr[i], arr[m-i] = arr[m-i], arr[i]
		}
		for i := 0; i < n/2; i++ {
			arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		}
		ans = append(ans, ind+1, n)
	}
	return
}
