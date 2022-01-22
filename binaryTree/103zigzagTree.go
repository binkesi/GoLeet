package binarytree

// https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int = make([][]int, 0)
	if root == nil {
		return res
	}
	var nodeQueue []*TreeNode = make([]*TreeNode, 0)
	var lev int = 0
	nodeQueue = append(nodeQueue, root)
	for {
		var childQueue []*TreeNode = make([]*TreeNode, 0)
		tmp := []int{}
		for ; len(nodeQueue) != 0; nodeQueue = nodeQueue[0 : len(nodeQueue)-1] {
			lenN := len(nodeQueue)
			if lev%2 == 0 {
				if nodeQueue[lenN-1].Left != nil {
					childQueue = append(childQueue, nodeQueue[lenN-1].Left)
				}
				if nodeQueue[lenN-1].Right != nil {
					childQueue = append(childQueue, nodeQueue[lenN-1].Right)
				}
			} else {
				if nodeQueue[lenN-1].Right != nil {
					childQueue = append(childQueue, nodeQueue[lenN-1].Right)
				}
				if nodeQueue[lenN-1].Left != nil {
					childQueue = append(childQueue, nodeQueue[lenN-1].Left)
				}
			}
			tmp = append(tmp, nodeQueue[lenN-1].Val)
		}
		res = append(res, tmp)
		if len(childQueue) == 0 {
			break
		} else {
			nodeQueue = childQueue
		}
		lev += 1
	}
	return res
}
