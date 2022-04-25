package hashmap

import "math/rand"

// https://leetcode-cn.com/problems/random-pick-index/submissions/

type Solution []int

func Constructor(nums []int) Solution {
	return nums
}

func (nums Solution) Pick(target int) (ans int) {
	cnt := 0
	for i, num := range nums {
		if num == target {
			cnt++ // 第 cnt 次遇到 target
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return
}
