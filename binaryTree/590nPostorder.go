package binarytree

// https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/

type NodeN struct {
	Val      int
	Children []*NodeN
}

func postorder(root *NodeN) (ans []int) {
	var post func(*NodeN)
	post = func(nn *NodeN) {
		if nn == nil {
			return
		}
		for _, v := range nn.Children {
			post(v)
		}
		ans = append(ans, nn.Val)
	}
	post(root)
	return
}

func postorder2(root *NodeN) (ans []int) {
	if root == nil {
		return
	}
	type Cnode struct {
		node *NodeN
		num  int
	}
	nStack := []Cnode{{root, 0}}
	for len(nStack) != 0 {
		n := len(nStack)
		top := nStack[n-1]
		nChild := len(top.node.Children)
		if nChild == 0 || top.num == nChild {
			ans = append(ans, top.node.Val)
			nStack = nStack[0 : n-1]
		} else {
			nStack[n-1].num++
			nStack = append(nStack, Cnode{top.node.Children[top.num], 0})
		}
	}
	return
}
