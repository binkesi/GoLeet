package treearray

import (
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

type MyCalendar struct {
	*redblacktree.Tree
}

func ConstructorCalendar() MyCalendar {
	t := redblacktree.NewWithIntComparator()
	t.Put(math.MaxInt32, nil)
	return MyCalendar{t}
}

func (this *MyCalendar) Book(start int, end int) bool {
	node, _ := this.Ceiling(end)
	it := this.IteratorAt(node)
	if !it.Prev() || it.Value().(int) <= start {
		this.Put(start, end)
		return true
	}
	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
