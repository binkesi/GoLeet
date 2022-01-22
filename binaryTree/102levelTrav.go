package binarytree

// https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type Queue struct {
	queue []interface{}
}

func newQueue(length int) *Queue {
	q := &Queue{
		queue: make([]interface{}, length),
	}
	return q
}

func (q *Queue) Inqueue(num interface{}) {
	if num == nil {
		return
	}
	q.queue = append(q.queue, num)
}

func (q *Queue) Dequeue() interface{} {
	if q.isEmpty() {
		return nil
	}
	tmp := q.queue[0]
	q.queue = q.queue[1:]
	return tmp
}

func (q *Queue) isEmpty() bool {
	return len(q.queue) == 0
}

func levelOrder(root *TreeNode) [][]int {
	var nodeQueue *Queue = newQueue(0)
	nodeQueue.Inqueue(root)
	var res [][]int = make([][]int, 0)
	for !nodeQueue.isEmpty() {
		tmp1 := make([]*TreeNode, 0)
		tmp2 := make([]int, 0)
		for !nodeQueue.isEmpty() {
			tmp1 = append(tmp1, nodeQueue.Dequeue().(*TreeNode))
		}
		for _, v := range tmp1 {
			if v != nil {
				nodeQueue.Inqueue(v.Left)
				nodeQueue.Inqueue(v.Right)
				tmp2 = append(tmp2, v.Val)
			}
		}
		if len(tmp2) != 0 {
			res = append(res, tmp2)
		}
	}
	return res
}
