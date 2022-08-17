package binarytree

// https://leetcode.cn/problems/deepest-leaves-sum/

func deepestLeavesSum(root *TreeNode) (ans int) {
	nodeQueue := []*TreeNode{root}
	for len(nodeQueue) != 0 {
		tmp := []*TreeNode{}
		sum := 0
		for i := range nodeQueue {
			p := nodeQueue[i]
			sum += p.Val
			if p.Left != nil {
				tmp = append(tmp, p.Left)
			}
			if p.Right != nil {
				tmp = append(tmp, p.Right)
			}
		}
		nodeQueue = tmp
		ans = sum
	}
	return
}
