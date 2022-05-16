package binarytree

// https://leetcode.cn/problems/successor-lcci/submissions/

func inorderSuccessor(root *TreeNode, p *TreeNode) (ans *TreeNode) {
	trav := root
	for trav != nil {
		if trav.Val > p.Val {
			ans = trav
			trav = trav.Left
		} else {
			trav = trav.Right
		}
	}
	return
}
