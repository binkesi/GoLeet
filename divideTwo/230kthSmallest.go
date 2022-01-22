package dividetwo

// https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var res int = 0
	inOrder(root, &k, &res)
	return res
}

func inOrder(node *TreeNode, k *int, val *int) {
	if node == nil {
		return
	}
	inOrder(node.Left, k, val)
	*k -= 1
	if *k == 0 {
		*val = node.Val
		return
	}
	inOrder(node.Right, k, val)
}
