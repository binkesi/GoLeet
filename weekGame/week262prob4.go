package weekgame

import "math"

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
	for i := 1; i <= n; i++ {
		for _, v := range ak[i-1] {
			for _, num := range a {
				ak[i] = append(ak[i], v+num)
			}
		}
	}
	for i := 1; i <= n; i++ {
		for _, v := range bk[i-1] {
			for _, num := range b {
				bk[i] = append(bk[i], v+num)
			}
		}
	}
	ans = math.MaxInt32
	for j := 0; j <= n; j++ {
		l, r := ak[j], bk[n-j]
		for _, v1 := range l {
			for _, v2 := range r {
				tmp := abs(sum - 2*(v1+v2))
				if tmp < ans {
					ans = tmp
				}
			}
		}
	}
	return
}
