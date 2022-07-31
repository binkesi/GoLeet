package binarytree

import "math"

// https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/submissions/

func maxLevelSum(root *TreeNode) int {
	ans, sum := 1, math.MinInt32
	curLev := []*TreeNode{root}
	nextLev := []*TreeNode{}
	lev := 1
	for {
		curSum := 0
		for len(curLev) != 0 {
			p := curLev[0]
			curLev = curLev[1:]
			curSum += p.Val
			if p.Left != nil {
				nextLev = append(nextLev, p.Left)
			}
			if p.Right != nil {
				nextLev = append(nextLev, p.Right)
			}
		}
		if curSum > sum {
			ans, sum = lev, curSum
		}
		if len(nextLev) == 0 {
			break
		}
		lev++
		curLev = nextLev
		nextLev = []*TreeNode{}
	}
	return ans
}
