package binarytree

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// 	Next  *Node
// }

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}
	for node := root; node.Left != nil; node = node.Left {
		tmp := node
		for ; tmp.Next != nil; tmp = tmp.Next {
			tmp.Left.Next = tmp.Right
			tmp.Right.Next = tmp.Next.Left
		}
		tmp.Left.Next = tmp.Right
	}
	return root
}
