package binarytree

// https://leetcode-cn.com/problems/binary-tree-pruning/

type TreeNodet struct {
	Val   int
	Left  *TreeNodet
	Right *TreeNodet
}

func pruneTree(root *TreeNodet) *TreeNodet {
	var isZero func(*TreeNodet) bool
	isZero = func(tn *TreeNodet) bool {
		if tn == nil {
			return true
		}
		return tn.Val == 0 && isZero(tn.Left) && isZero(tn.Right)
	}
	var removeZero func(*TreeNodet) *TreeNodet
	removeZero = func(tn *TreeNodet) *TreeNodet {
		if isZero(tn) {
			return nil
		}
		if isZero(tn.Left) {
			tn.Left = nil
		} else {
			removeZero(tn.Left)
		}
		if isZero(tn.Right) {
			tn.Right = nil
		} else {
			removeZero(tn.Right)
		}
		return tn
	}
	return removeZero(root)
}
