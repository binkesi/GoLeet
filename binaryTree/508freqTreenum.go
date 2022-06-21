package binarytree

// https://leetcode.cn/problems/most-frequent-subtree-sum/submissions/

func findFrequentTreeSum(root *TreeNode) (ans []int) {
	var calSum func(node *TreeNode) int
	vMap := make(map[int]int)
	calSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		node.Val = node.Val + calSum(node.Left) + calSum(node.Right)
		val := node.Val
		if _, ok := vMap[val]; !ok {
			vMap[val] = 1
		} else {
			vMap[val] += 1
		}
		return val
	}
	_ = calSum(root)
	times := 0
	for k, v := range vMap {
		if v > times {
			ans = []int{}
			ans = append(ans, k)
			times = v
		} else if v == times {
			ans = append(ans, k)
		}
	}
	return
}
