package binarytree

// https://leetcode.cn/problems/longest-univalue-path/

func longestUnivaluePath(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		curLeft, curRight := 0, 0
		if node.Left == nil && node.Right == nil {
			return 0
		}
		if node.Left != nil {
			lvalue := dfs(node.Left)
			if node.Val == node.Left.Val {
				curLeft = curLeft + lvalue + 1
			}
		}
		if node.Right != nil {
			rvalue := dfs(node.Right)
			if node.Val == node.Right.Val {
				curRight = curRight + rvalue + 1
			}
		}
		if curLeft+curRight > ans {
			ans = curLeft + curRight
		}
		return max(curLeft, curRight)
	}
	dfs(root)
	return
}
