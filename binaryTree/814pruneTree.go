package binarytree

// https://leetcode-cn.com/problems/binary-tree-pruning/

func pruneTree(root *TreeNode) *TreeNode {
	var removeZero func(node *TreeNode) bool
	removeZero = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if removeZero(node.Left) {
			node.Left = nil
		}
		if removeZero(node.Right) {
			node.Right = nil
		}
		if node.Val == 1 {
			return false
		}
		if node.Val == 0 && node.Left == nil && node.Right == nil {
			return true
		}
		return false
	}
	if removeZero(root) {
		return nil
	}
	return root
}
