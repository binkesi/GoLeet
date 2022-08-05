package binarytree

// https://leetcode.cn/problems/add-one-row-to-tree/

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		ans := TreeNode{Val: val}
		ans.Left = root
		return &ans
	}
	nodeQ := []*TreeNode{root}
	for depth != 2 {
		tmp := []*TreeNode{}
		for len(nodeQ) != 0 {
			p := nodeQ[0]
			nodeQ = nodeQ[1:]
			if p.Left != nil {
				tmp = append(tmp, p.Left)
			}
			if p.Right != nil {
				tmp = append(tmp, p.Right)
			}
		}
		depth--
		nodeQ = tmp
	}
	l := len(nodeQ)
	for i := 0; i < l; i++ {
		p := nodeQ[i]
		if p.Left != nil {
			node := TreeNode{Val: val}
			node.Left = p.Left
			p.Left = &node
		} else {
			node := TreeNode{Val: val}
			p.Left = &node
		}
		if p.Right != nil {
			node := TreeNode{Val: val}
			node.Right = p.Right
			p.Right = &node
		} else {
			node := TreeNode{Val: val}
			p.Right = &node
		}
	}
	return root
}
