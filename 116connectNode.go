package goleet

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	var nodePar []*Node = make([]*Node, 0)
	var nodeChild []*Node = make([]*Node, 0)
	if root == nil {
		return root
	}
	root.Next = nil
	nodePar = append(nodePar, root)
	for len(nodePar) != 0 {
		for len(nodePar) != 0 {
			if len(nodePar) == 1 {
				nodePar[0].Next = nil
			} else {
				nodePar[0].Next = nodePar[1]
			}
			if nodePar[0].Left != nil {
				nodeChild = append(nodeChild, nodePar[0].Left, nodePar[0].Right)
			}
			nodePar = nodePar[1:]
		}
		nodePar, nodeChild = nodeChild, nodePar
		//fmt.Println(nodePar, nodeChild)
	}
	return root
}
