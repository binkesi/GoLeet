package dfsbfs

import "math"

// https://leetcode.cn/problems/find-largest-value-in-each-tree-row/submissions/

func largestValues(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	ans = append(ans, root.Val)
	tmp := []*TreeNode{}
	if root.Left != nil {
		tmp = append(tmp, root.Left)
	}
	if root.Right != nil {
		tmp = append(tmp, root.Right)
	}
	for len(tmp) != 0 {
		queue = tmp
		maxNum := math.MinInt32
		tmp = []*TreeNode{}
		n := len(queue)
		for i := 0; i < n; i++ {
			p := queue[i]
			if p.Val > maxNum {
				maxNum = p.Val
			}
			if p.Left != nil {
				tmp = append(tmp, p.Left)
			}
			if p.Right != nil {
				tmp = append(tmp, p.Right)
			}
		}
		ans = append(ans, maxNum)
	}
	return
}
