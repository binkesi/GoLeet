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
