package goleet

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxPath := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		left := max(maxGain(tn.Left), 0)
		right := max(maxGain(tn.Right), 0)
		maxRoad := tn.Val + left + right
		maxPath = max(maxPath, maxRoad)
		return tn.Val + max(left, right)
	}
	maxGain(root)
	return maxPath
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
