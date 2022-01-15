package goleet

// https://leetcode-cn.com/problems/sort-list/

type ListNodel struct {
	Val  int
	Next *ListNodel
}

func mergeTwoList(lis1, lis2 *ListNodel) *ListNodel {
	dummyHead := &ListNodel{}
	temp, temp1, temp2 := dummyHead, lis1, lis2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}

func mergeSortList(head, tail *ListNodel) *ListNodel {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return mergeTwoList(mergeSortList(head, mid), mergeSortList(mid, tail))
}

func sortList(head *ListNodel) *ListNodel {
	return mergeSortList(head, nil)
}
