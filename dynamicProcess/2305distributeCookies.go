package dynamicprocess

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/fair-distribution-of-cookies/submissions/

func distributeCookies(cookies []int, k int) int {
	ans := math.MaxInt32
	sort.Slice(cookies, func(i, j int) bool { return cookies[i] > cookies[j] })
	up, sum, l := 0, 0, len(cookies)
	for i := 0; i < l; i++ {
		sum += cookies[i]
		if i == l/k {
			up = sum
		}
	}
	dp := make([]int, k)
	var put func(arr []int, ind int)
	put = func(arr []int, ind int) {
		// 第一包分给谁都一样，所以分配给第一个小朋友，减少回溯分支
		if ind == 0 {
			arr[0] = cookies[0]
			put(arr, 1)
			return
		}
		cnt := 0
		for i := 0; i < k; i++ {
			// 如果某位小朋友的饼干数量比当前的答案还多，显然继续回溯下去也无法成为最优答案，直接返回。
			if arr[i] > ans {
				return
			}
			if arr[i] == 0 {
				cnt++
			}
			// 如果当前空手的小朋友比剩余的饼干包还多，显然无法得出最优解，直接返回
			if cnt > l-ind {
				return
			}
		}
		res := maxN(arr)
		if res >= up {
			return
		}
		if ind == l {
			if res < ans {
				ans = res
			}
		} else {
			ele := cookies[ind]
			for i := range arr {
				nextArr := make([]int, k)
				copy(nextArr, arr)
				if nextArr[i]+ele < up {
					nextArr[i] += ele
					put(nextArr, ind+1)
				}
			}
		}
	}
	put(dp, 0)
	return ans
}

func maxN(arr []int) int {
	l, ans := len(arr), 0
	for i := 0; i < l; i++ {
		if arr[i] > ans {
			ans = arr[i]
		}
	}
	return ans
}
