package goleet

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	lenL := len(lists)
	if lenL == 0 {
		return res
	}
	if lenL == 1 {
		res = lists[0]
		return res
	}
	res = lists[0]
	for i := 1; i < lenL; i++ {
		res = mertwo(res, lists[i])
	}
	return res
}

func mertwo(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res *ListNode = &ListNode{}
	tmp := res
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tmp.Next = l1
			l1 = l1.Next
		} else {
			tmp.Next = l2
			l2 = l2.Next
		}
		tmp = tmp.Next
	}
	if l1 == nil {
		tmp.Next = l2
	} else if l2 == nil {
		tmp.Next = l1
	}
	return res.Next
}
