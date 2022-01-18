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
		res = mergeTwoList(res, lists[i])
	}
	return res
}

func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode = &ListNode{}

}
