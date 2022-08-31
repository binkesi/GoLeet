package binarytree

// https://leetcode.cn/problems/maximum-binary-tree-ii/

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	node := TreeNode{Val: val}
	var pre *TreeNode
	pre = nil
	cur := root
	for cur != nil && cur.Val > val {
		pre = cur
		cur = cur.Right
	}
	if pre == nil {
		node.Left = cur
		return &node
	}
	pre.Right = &node
	node.Left = cur
	return root
}
