package binarytree

import "fmt"

// https://leetcode.cn/problems/find-duplicate-subtrees/

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	repeat := make(map[*TreeNode]struct{})
	seen := make(map[string]*TreeNode)
	var dfs func(*TreeNode) string
	dfs = func(tn *TreeNode) string {
		if tn == nil {
			return ""
		}
		serial := fmt.Sprintf("%d(%s)(%s)", tn.Val, dfs(tn.Left), dfs(tn.Right))
		if n, ok := seen[serial]; !ok {
			seen[serial] = tn
		} else {
			repeat[n] = struct{}{}
		}
		return serial
	}
	dfs(root)
	ans := make([]*TreeNode, 0, len(repeat))
	for node := range repeat {
		ans = append(ans, node)
	}
	return ans
}
