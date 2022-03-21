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

func findTarget2(root *TreeNode, k int) bool {
	nArr := []int{}
	var inorder func(*TreeNode)
	inorder = func(tn *TreeNode) {
		if tn == nil {
			return
		}
		inorder(tn.Left)
		nArr = append(nArr, tn.Val)
		inorder(tn.Right)
	}
	inorder(root)
	p1, p2 := 0, len(nArr)-1
	for p1 < p2 {
		sum := nArr[p1] + nArr[p2]
		if sum == k {
			return true
		}
		if sum < k {
			p1++
		} else {
			p2--
		}
	}
	return false
}
