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

func preorder2(root *Nnode) (ans []int) {
	if root == nil {
		return
	}
	type nInfo struct {
		node *Nnode
		cnt  int
	}
	nodes := []nInfo{}
	nodes = append(nodes, nInfo{root, 0})
	for len(nodes) != 0 {
		n := len(nodes)
		top := nodes[n-1]
		if len(top.node.Children) == top.cnt {
			if top.cnt == 0 {
				ans = append(ans, top.node.Val)
			}
			nodes = nodes[0 : n-1]
		} else {
			if top.cnt == 0 {
				ans = append(ans, top.node.Val)
			}
			nodes = append(nodes, nInfo{top.node.Children[top.cnt], 0})
			nodes[n-1].cnt += 1
		}
	}
	return
}
