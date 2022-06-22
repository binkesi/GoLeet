package binarytree

// https://leetcode.cn/problems/find-bottom-left-tree-value/submissions/

type NodeLev struct {
	node *TreeNode
	lev  int
}

func findBottomLeftValue(root *TreeNode) (ans int) {
	queue := []NodeLev{{node: root, lev: 1}}
	maxLev := 1
	ans = root.Val
	for len(queue) != 0 {
		p := queue[0]
		queue = queue[1:]
		if p.node.Left != nil {
			queue = append(queue, NodeLev{node: p.node.Left, lev: p.lev + 1})
			if p.lev+1 > maxLev {
				maxLev = p.lev + 1
				ans = p.node.Left.Val
			}
		}
		if p.node.Right != nil {
			queue = append(queue, NodeLev{node: p.node.Right, lev: p.lev + 1})
			if p.lev+1 > maxLev {
				maxLev = p.lev + 1
				ans = p.node.Right.Val
			}
		}
	}
	return
}
