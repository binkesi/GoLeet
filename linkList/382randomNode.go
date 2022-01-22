package linklist

import (
	"math/rand"
)

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() int {
	if this.head == nil {
		return 0
	}
	slow, fast := this.head, this.head
	n := 0
	for fast != nil {
		n += 1
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			n += 1
			fast = fast.Next
		}
	}
	rand.Seed(2)
	randInt := rand.Intn(n)
	node := this.head
	for i := 1; i <= randInt; i++ {
		node = node.Next
	}
	return node.Val
}
