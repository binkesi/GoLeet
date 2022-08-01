package binarytree

// https://leetcode.cn/problems/logical-or-of-two-binary-grids-represented-as-quad-trees/submissions/

type QuadNode struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *QuadNode
	TopRight    *QuadNode
	BottomLeft  *QuadNode
	BottomRight *QuadNode
}

func intersect(quadTree1, quadTree2 *QuadNode) *QuadNode {
	if quadTree1.IsLeaf {
		if quadTree1.Val {
			return &QuadNode{Val: true, IsLeaf: true}
		}
		return quadTree2
	}
	if quadTree2.IsLeaf {
		return intersect(quadTree2, quadTree1)
	}
	o1 := intersect(quadTree1.TopLeft, quadTree2.TopLeft)
	o2 := intersect(quadTree1.TopRight, quadTree2.TopRight)
	o3 := intersect(quadTree1.BottomLeft, quadTree2.BottomLeft)
	o4 := intersect(quadTree1.BottomRight, quadTree2.BottomRight)
	if o1.IsLeaf && o2.IsLeaf && o3.IsLeaf && o4.IsLeaf && o1.Val == o2.Val && o1.Val == o3.Val && o1.Val == o4.Val {
		return &QuadNode{Val: o1.Val, IsLeaf: true}
	}
	return &QuadNode{false, false, o1, o2, o3, o4}
}
