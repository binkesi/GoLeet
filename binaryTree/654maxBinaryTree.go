package binarytree

// https://leetcode.cn/problems/maximum-binary-tree/submissions/

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxNum := nums[0]
	ind := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
			ind = i
		}
	}
	node := TreeNode{Val: maxNum}
	node.Left = constructMaximumBinaryTree(nums[0:ind])
	node.Right = constructMaximumBinaryTree(nums[ind+1:])
	return &node
}
