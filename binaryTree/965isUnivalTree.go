package binarytree

// https://leetcode.cn/problems/univalued-binary-tree/submissions/

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	var preOrder func(*TreeNode) bool
	preOrder = func(t *TreeNode) bool {
		if t.Val != val {
			return false
		}
		return (t.Left == nil || preOrder(t.Left)) && (t.Right == nil || preOrder(t.Right))
	}
	return preOrder(root)
}
