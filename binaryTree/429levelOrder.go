package binarytree

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

func levelOrderN(root *NodeN) (ans [][]int) {
	if root == nil {
		return
	}
	queue := []*NodeN{root}
	ans = append(ans, []int{root.Val})
	for len(queue) != 0 {
		nextLevel := []*NodeN{}
		tmp := []int{}
		for _, node := range queue {
			nextLevel = append(nextLevel, node.Children...)
			for _, child := range node.Children {
				tmp = append(tmp, child.Val)
			}
		}
		if len(tmp) != 0 {
			ans = append(ans, tmp)
		}
		queue = nextLevel
	}
	return
}
