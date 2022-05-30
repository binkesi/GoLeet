package dfsbfs

// https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/submissions/

func sumRootToLeaf(root *TreeNode) (ans int) {
	var dfs func(node *TreeNode, cur int)
	dfs = func(node *TreeNode, cur int) {
		if node.Left == nil && node.Right == nil {
			ans += cur*2 + node.Val
		}
		if node.Left != nil {
			dfs(node.Left, cur*2+node.Val)
		}
		if node.Right != nil {
			dfs(node.Right, cur*2+node.Val)
		}
	}
	dfs(root, 0)
	return
}
