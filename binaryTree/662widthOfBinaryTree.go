package binarytree

// https://leetcode.cn/problems/maximum-width-of-binary-tree/submissions/

type nodeNum struct {
	node *TreeNode
	num  int
}

func widthOfBinaryTree(root *TreeNode) int {
	myQue := []nodeNum{{node: root, num: 1}}
	maxLen := 1
	for len(myQue) != 0 {
		tmp := []nodeNum{}
		for _, v := range myQue {
			if v.node == nil {
				tmp = append(tmp, nodeNum{node: nil, num: 2 * v.num})
			} else {
				tmp = append(tmp, nodeNum{node: v.node.Left, num: 1}, nodeNum{node: v.node.Right, num: 1})
			}
		}
		le, ri := 0, len(tmp)-1
		for le < len(tmp) && tmp[le].node == nil {
			le++
		}
		for ri >= 0 && tmp[ri].node == nil {
			ri--
		}
		if le > ri {
			myQue = []nodeNum{}
		} else {
			myQue = tmp[le : ri+1]
			curLen := 0
			for _, v := range myQue {
				curLen += v.num
			}
			maxLen = max(maxLen, curLen)
		}
	}
	return maxLen
}

func widthOfBinaryTree2(root *TreeNode) int {
	maxWid := 1
	leftMap := make(map[int]int)
	leftMap[1] = 1
	var dfs func(node *TreeNode, depth, ind int)
	dfs = func(node *TreeNode, depth, ind int) {
		if node == nil {
			return
		}
		if _, ok := leftMap[depth]; !ok {
			leftMap[depth] = ind
		}
		maxWid = max(maxWid, ind-leftMap[depth]+1)
		dfs(node.Left, depth+1, ind<<1)
		dfs(node.Right, depth+1, ind<<1|1)
	}
	dfs(root, 1, 0)
	return maxWid
}
