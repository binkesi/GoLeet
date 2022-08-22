package binarytree

import (
	"math"
	"strconv"
)

// https://leetcode.cn/problems/print-binary-tree/

func printTree(root *TreeNode) [][]string {
	var calHeight func(node *TreeNode) int
	type calNode struct {
		node *TreeNode
		x    int
		y    int
	}
	calHeight = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		return max(calHeight(node.Left), calHeight(node.Right)) + 1
	}
	height := calHeight(root)
	res := make([][]string, height+1)
	n := int(math.Pow(2.0, float64(height+1)) - 1)
	for i := range res {
		res[i] = make([]string, n)
	}
	res[0][(n-1)/2] = strconv.Itoa(root.Val)
	que := []calNode{{root, 0, (n - 1) / 2}}
	for len(que) != 0 {
		tmp := make([]calNode, 0)
		for len(que) != 0 {
			p := que[0]
			que = que[1:]
			r := p.x
			dif := int(math.Pow(2.0, float64(height-r-1)))
			if p.node.Left != nil {
				tmp = append(tmp, calNode{node: p.node.Left, x: p.x + 1, y: p.y - dif})
				res[p.x+1][p.y-dif] = strconv.Itoa(p.node.Left.Val)
			}
			if p.node.Right != nil {
				tmp = append(tmp, calNode{node: p.node.Right, x: p.x + 1, y: p.y + dif})
				res[p.x+1][p.y+dif] = strconv.Itoa(p.node.Right.Val)
			}
		}
		que = tmp
	}
	return res
}
