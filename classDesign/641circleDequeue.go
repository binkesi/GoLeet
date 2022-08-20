package classdesign

// https://leetcode.cn/problems/design-circular-deque/

type MyCircularDeque struct {
	myQueue [4000]int
	head    int
	tail    int
	cap     int
}

func ConstructorDequeue(k int) MyCircularDeque {
	return MyCircularDeque{[4000]int{}, 2000, 1999, k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.myQueue[this.head-1] = value
	this.head--
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.myQueue[this.tail+1] = value
	this.tail++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head++
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.tail--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.myQueue[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.myQueue[this.tail]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.head-1 == this.tail
}

func (this *MyCircularDeque) IsFull() bool {
	return this.tail-this.head == this.cap-1
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
