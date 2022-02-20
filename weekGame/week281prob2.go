package weekgame

// https://leetcode-cn.com/contest/weekly-contest-281/problems/merge-nodes-in-between-zeros/

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	ptr, h := head.Next, head.Next
	for ptr != nil {
		sum := 0
		for ptr.Val != 0 {
			sum += ptr.Val
			ptr = ptr.Next
		}
		h.Next = ptr.Next
		h.Val = sum
		h = h.Next
		ptr = ptr.Next
	}
	head = head.Next
	return head
}
