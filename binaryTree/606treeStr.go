package binarytree

import "strconv"

// https://leetcode-cn.com/problems/construct-string-from-binary-tree/

type TreeNodeS struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) (ans string) {
	var preOrder func(*TreeNode)
	preOrder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		ans += strconv.Itoa(tn.Val)
		if tn.Left == nil && tn.Right == nil {
			return
		}
		ans += "("
		preOrder(tn.Left)
		ans += ")"
		if tn.Right != nil {
			ans += "("
			preOrder(tn.Right)
			ans += ")"
		}
	}
	preOrder(root)
	return
}
