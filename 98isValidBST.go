package goleet

import "math"

// https://leetcode-cn.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	var res int = math.MinInt64
	var inorder func(node *TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		res1 := inorder(node.Left)
		tmp := node.Val
		if tmp <= res {
			return false
		} else {
			res = tmp
		}
		res2 := inorder(node.Right)
		return res1 && res2
	}
	return inorder(root)
}
