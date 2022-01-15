package goleet

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var root *TreeNode = &TreeNode{
		Val: preorder[0],
	}
	i := 0
	for ; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[0:i])+1], inorder[0:i])
	root.Right = buildTree(preorder[len(inorder[0:i])+1:], inorder[i+1:])
	return root
}
