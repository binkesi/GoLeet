package weekgame

// https://leetcode-cn.com/problems/create-binary-tree-from-descriptions/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	nMap := make(map[int]struct{})
	tMap := make(map[int]map[int]int)
	root := TreeNode{}
	for _, v := range descriptions {
		par, child, dir := v[0], v[1], v[2]
		nMap[par] = struct{}{}
		if _, ok := tMap[par]; !ok {
			tMap[par] = make(map[int]int)
			tMap[par][child] = dir
		} else {
			tMap[par][child] = dir
		}
	}
	for _, v := range descriptions {
		delete(nMap, v[1])
	}
	for k := range nMap {
		root.Val = k
	}
	var front func(*TreeNode)
	front = func(tn *TreeNode) {
		if _, ok := tMap[tn.Val]; !ok {
			return
		} else {
			for k, v := range tMap[tn.Val] {
				if v == 1 {
					left := TreeNode{Val: k}
					tn.Left = &left
					front(tn.Left)
				}
				if v == 0 {
					right := TreeNode{Val: k}
					tn.Right = &right
					front(tn.Right)
				}
			}
		}
	}
	front(&root)
	return &root
}
