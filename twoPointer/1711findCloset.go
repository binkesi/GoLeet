package twopointer

// https://leetcode.cn/problems/find-closest-lcci/submissions/

func findClosest(words []string, word1 string, word2 string) (ans int) {
	arr1, arr2 := []int{}, []int{}
	for i, w := range words {
		if w == word1 {
			arr1 = append(arr1, i)
		}
		if w == word2 {
			arr2 = append(arr2, i)
		}
	}
	m, n := len(arr1), len(arr2)
	ans = abs(arr1[0] - arr2[0])
	for i, j := 0, 0; i < m && j < n; {
		if abs(arr1[i]-arr2[j]) < ans {
			ans = abs(arr1[i] - arr2[j])
		}
		if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}
	return
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}
