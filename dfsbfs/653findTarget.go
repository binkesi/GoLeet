package dfsbfs

// https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/submissions/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var dfs func(*TreeNode) bool
	nMap := make(map[int]struct{})
	dfs = func(tn *TreeNode) bool {
		if tn == nil {
			return false
		}
		if _, ok := nMap[k-tn.Val]; ok {
			return true
		} else {
			nMap[tn.Val] = struct{}{}
			return dfs(tn.Left) || dfs(tn.Right)
		}
	}
	return dfs(root)
}
