package binarytree

// https://leetcode.cn/problems/complete-binary-tree-inserter/

type CBTInserter struct {
	root      *TreeNode
	nodeQueue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	return CBTInserter{root: root, nodeQueue: []*TreeNode{root}}
}

func (this *CBTInserter) Insert(val int) int {
	q := this.nodeQueue
	for {
		p := q[0]
		q = q[1:]
		if p.Left == nil {
			p.Left = &TreeNode{Val: val}
			return p.Val
		}
		if p.Right == nil {
			p.Right = &TreeNode{Val: val}
			return p.Val
		}
		q = append(q, p.Left)
		q = append(q, p.Right)
	}
}

func (this *CBTInserter) Get_root() *TreeNode {
	return this.root
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(val);
 * param_2 := obj.Get_root();
 */
