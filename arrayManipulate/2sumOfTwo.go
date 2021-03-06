// https://leetcode-cn.com/problems/add-two-numbers/
package arraymanipulate

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	}
	if l2.Val == 0 && l2.Next == nil {
		return l1
	}
	l3 := ListNode{}
	p := &l3
	res := p
	carry := 0
	for {
		sum := ((*l1).Val + (*l2).Val + carry) % 10
		carry = ((*l1).Val + (*l2).Val + carry) / 10
		(*p).Val = sum
		l1 = (*l1).Next
		l2 = (*l2).Next
		if l1 == nil || l2 == nil {
			break
		}
		(*p).Next = &ListNode{}
		p = (*p).Next
	}
	if l1 == nil {
		(*p).Next = l2
	} else {
		(*p).Next = l1
	}
	p = (*p).Next
	for {
		if p == nil || carry == 0 {
			break
		}
		c := ((*p).Val + carry) / 10
		(*p).Val = ((*p).Val + carry) % 10
		carry = c
		p = (*p).Next
	}
	if carry == 1 {
		pt := res
		for {
			if (*pt).Next == nil {
				(*pt).Next = &ListNode{}
				(*((*pt).Next)).Val = 1
				break
			}
			pt = (*pt).Next
		}
	}
	return res
}
