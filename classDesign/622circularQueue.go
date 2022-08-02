package classdesign

// https://leetcode.cn/problems/design-circular-queue/

type MyCircularQueue struct {
	l    int
	head int
	tail int
	qu   []int
}

func ConstructorCir(k int) MyCircularQueue {
	return MyCircularQueue{
		l:    k + 1,
		head: 0,
		tail: 0,
		qu:   make([]int, k+1),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.qu[this.tail] = value
		this.tail = (this.tail + 1) % this.l
		return true
	}
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.head = (this.head + 1) % this.l
		return true
	}
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.qu[this.head]
	}
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.qu[(this.tail+this.l-1)%this.l]
	}
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%this.l == this.head
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
