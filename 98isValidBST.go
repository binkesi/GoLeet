package goleet

// https://leetcode-cn.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	var res []int = make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	for i := 1; i < len(res); i++ {
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}
