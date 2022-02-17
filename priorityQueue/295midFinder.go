package priorityqueue

import "container/heap"

// https://leetcode-cn.com/problems/find-median-from-data-stream/

type MedianFinder struct {
	sh sHeap
	bh bHeap
}

func Constructor() MedianFinder {
	sh, bh := sHeap{}, bHeap{}
	heap.Init(&sh)
	heap.Init(&bh)
	return MedianFinder{sh: sh, bh: bh}
}

func (this *MedianFinder) AddNum(num int) {
	l1, l2 := this.sh.Len(), this.bh.Len()
	if l1 == 0 {
		heap.Push(&this.sh, num)
	} else if l2 == 0 {
		heap.Push(&this.bh, num)
	}
	if l1 != 0 && l2 != 0 {
		l, r := this.sh[0], this.bh[0]
		if l < r {
			heap.Pop(&this.sh)
			heap.Pop(&this.bh)
			heap.Push(&this.sh, r)
			heap.Push(&this.bh, l)
		}
		if num >= r {
			heap.Push(&this.sh, num)
		} else {
			heap.Push(&this.bh, num)
		}
		if this.sh.Len() > this.bh.Len()+1 {
			a := this.sh[0]
			heap.Pop(&this.sh)
			heap.Push(&this.bh, a)
		} else if this.sh.Len() < this.bh.Len() {
			b := this.bh[0]
			heap.Pop(&this.bh)
			heap.Push(&this.sh, b)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.sh.Len() == this.bh.Len() {
		return float64(this.sh[0]+this.bh[0]) / 2
	} else {
		return float64(this.sh[0])
	}
}

type sHeap []int

func (h sHeap) Len() int            { return len(h) }
func (h sHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h sHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h *sHeap) Push(a interface{}) { *h = append(*h, a.(int)) }
func (h *sHeap) Pop() interface{} {
	old := *h
	n := h.Len()
	res := old[n-1]
	*h = old[0 : n-1]
	return res
}

type bHeap []int

func (h bHeap) Len() int            { return len(h) }
func (h bHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h bHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h *bHeap) Push(a interface{}) { *h = append(*h, a.(int)) }
func (h *bHeap) Pop() interface{} {
	old := *h
	n := h.Len()
	res := old[n-1]
	*h = old[0 : n-1]
	return res
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
