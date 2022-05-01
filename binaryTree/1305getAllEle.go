package binarytree

// https://leetcode.cn/problems/all-elements-in-two-binary-search-trees/submissions/

func getAllElements(root1 *TreeNode, root2 *TreeNode) (ans []int) {
	lis1, lis2 := []int{}, []int{}
	var midOrder func(res *[]int, root *TreeNode)
	midOrder = func(lis *[]int, root *TreeNode) {
		if root == nil {
			return
		}
		midOrder(lis, root.Left)
		*lis = append(*lis, root.Val)
		midOrder(lis, root.Right)
	}
	midOrder(&lis1, root1)
	midOrder(&lis2, root2)
	l1, l2 := len(lis1), len(lis2)
	p1, p2 := 0, 0
	for p1 < l1 && p2 < l2 {
		if lis1[p1] < lis2[p2] {
			ans = append(ans, lis1[p1])
			p1++
		} else {
			ans = append(ans, lis2[p2])
			p2++
		}
	}
	if p1 != l1 {
		ans = append(ans, lis1[p1:]...)
	} else {
		ans = append(ans, lis2[p2:]...)
	}
	return
}
