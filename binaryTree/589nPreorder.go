package binarytree

type Nnode struct {
	Val      int
	Children []*Nnode
}

// https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/

func preorder(root *Nnode) (ans []int) {
	var porder func(*Nnode)
	porder = func(n *Nnode) {
		if n == nil {
			return
		}
		ans = append(ans, n.Val)
		for _, v := range n.Children {
			porder(v)
		}
	}
	porder(root)
	return
}
