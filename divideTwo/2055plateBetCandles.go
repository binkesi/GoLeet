package dividetwo

// https://leetcode-cn.com/problems/plates-between-candles/

func platesBetweenCandles(s string, queries [][]int) (ans []int) {
	n := len(s)
	tmp := make([]int, n)
	tmp[0] = 0
	for i := 1; i < n; i++ {
		if s[i] == '|' {
			tmp[i] = tmp[i-1]
		} else {
			tmp[i] = tmp[i-1] + 1
		}
	}
	left, right := make([]int, n), make([]int, n)
	if s[0] == '|' {
		left[0] = 0
	} else {
		left[0] = -1
	}
	if s[n-1] == '|' {
		right[n-1] = n - 1
	} else {
		right[n-1] = -1
	}
	for i := 1; i < n; i++ {
		if s[i] == '|' {
			left[i] = i
		} else {
			left[i] = left[i-1]
		}
	}
	for j := n - 2; j >= 0; j-- {
		if s[j] == '|' {
			right[j] = j
		} else {
			right[j] = right[j+1]
		}
	}
	for _, v := range queries {
		l, r := v[0], v[1]
		if right[l] == -1 || left[r] == -1 || left[r] <= right[l] {
			ans = append(ans, 0)
			continue
		} else {
			res := tmp[left[r]] - tmp[right[l]]
			ans = append(ans, res)
		}
	}
	return
}
