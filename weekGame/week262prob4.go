package weekgame

import (
	"math"
	"math/bits"
	"sort"
)

// https://leetcode-cn.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/

func minimumDifference2(nums []int) (ans int) {
	n := len(nums) / 2
	sum := 0
	for k := 0; k < len(nums); k++ {
		sum += nums[k]
	}
	a, b := nums[0:n], nums[n:]
	ak, bk := make([][]int, n+1), make([][]int, n+1)
	ak[0], bk[0] = []int{0}, []int{0}
	stats := 1 << n
	for i := 0; i < stats; i++ {
		tmp := i
		cnt := bits.OnesCount(uint(tmp))
		resA, resB := 0, 0
		for j := n - 1; tmp > 0; j-- {
			if tmp&1 == 1 {
				resA += a[j]
				resB += b[j]
			}
			tmp = tmp >> 1
		}
		ak[cnt] = append(ak[cnt], resA)
		bk[cnt] = append(bk[cnt], resB)
	}
	ans = math.MaxInt32
	for j := 0; j <= n; j++ {
		l, r := ak[j], bk[n-j]
		sort.Ints(r)
		lr := len(r)
		for _, v1 := range l {
			left, right := 0, lr-1
			for left < right-1 {
				mid := left + (right-left)/2
				tmp := sum - 2*(v1+r[mid])
				if tmp == 0 {
					ans = 0
					return
				}
				if tmp > 0 {
					left = mid
				} else if tmp < 0 {
					right = mid
				}
			}
			ans = minNum(ans, abs(sum-2*(v1+r[left])), abs(sum-2*(v1+r[right])))
		}
	}
	return
}

func minNum(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}
