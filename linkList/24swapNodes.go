// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
package linklist

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func swapPairs(head *ListNode) *ListNode {
	if head == nil || (*head).Next == nil {
		return head
	}
	p1, p2 := head, (*head).Next
	(*p1).Next = (*p2).Next
	(*p2).Next = head
	head = p2
	tmp := p1
	p1 = (*p1).Next
	if p1 == nil {
		return head
	}
	p2 = (*p1).Next
	for {
		if p1 == nil {
			break
		}
		p2 = (*p1).Next
		if p2 == nil {
			break
		}
		(*p1).Next = (*p2).Next
		(*p2).Next = p1
		(*tmp).Next = p2
		tmp = p1
		p1 = (*p1).Next
	}
	return head
}
