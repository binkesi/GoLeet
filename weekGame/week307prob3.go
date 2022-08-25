package weekgame

// https://leetcode.cn/contest/weekly-contest-307/problems/amount-of-time-for-binary-tree-to-be-infected/

func amountOfTime(root *TreeNode, start int) int {
	nMap := make(map[int][]int)
	var preOrder func(node *TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		if _, ok := nMap[node.Val]; !ok {
			nMap[node.Val] = make([]int, 0)
		}
		if node.Left != nil {
			nMap[node.Val] = append(nMap[node.Val], node.Left.Val)
			if _, ok := nMap[node.Left.Val]; !ok {
				nMap[node.Left.Val] = make([]int, 0)
			}
			nMap[node.Left.Val] = append(nMap[node.Left.Val], node.Val)
		}
		if node.Right != nil {
			nMap[node.Val] = append(nMap[node.Val], node.Right.Val)
			if _, ok := nMap[node.Right.Val]; !ok {
				nMap[node.Right.Val] = make([]int, 0)
			}
			nMap[node.Right.Val] = append(nMap[node.Right.Val], node.Val)
		}
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	visited := make(map[int]struct{})
	visited[start] = struct{}{}
	maxMin := 0
	myQue := [][2]int{{start, 0}}
	for len(myQue) != 0 {
		p := myQue[0]
		if maxMin < p[1] {
			maxMin = p[1]
		}
		myQue = myQue[1:]
		for _, v := range nMap[p[0]] {
			if _, ok := visited[v]; !ok {
				visited[v] = struct{}{}
				myQue = append(myQue, [2]int{v, p[1] + 1})
			}
		}
	}
	return maxMin
}
