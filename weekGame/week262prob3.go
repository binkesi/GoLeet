package weekgame

import "container/heap"

type StockPrice struct {
	lHeap, sHeap hp
	price        map[int]int
	time         int
}

func Constructor22() StockPrice {
	return StockPrice{
		price: make(map[int]int),
	}
}

func (this *StockPrice) Update(timestamp int, price int) {
	heap.Push(&this.lHeap, pairS{-price, timestamp})
	heap.Push(&this.sHeap, pairS{price, timestamp})
	this.price[timestamp] = price
	if timestamp > this.time {
		this.time = timestamp
	}
}

func (this *StockPrice) Current() int {
	return this.price[this.time]
}

func (this *StockPrice) Maximum() int {
	for {
		if p := this.lHeap[0]; -p.price == this.price[p.time] {
			return -p.price
		}
		heap.Pop(&this.lHeap)
	}
}

func (this *StockPrice) Minimum() int {
	for {
		if p := this.sHeap[0]; p.price == this.price[p.time] {
			return p.price
		}
		heap.Pop(&this.sHeap)
	}
}

type pairS struct {
	price int
	time  int
}

type hp []pairS

func (h hp) Len() int { return len(h) }
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h hp) Less(i, j int) bool  { return h[i].price < h[j].price }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pairS)) }
func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

/**
 * Your StockPrice object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Update(timestamp,price);
 * param_2 := obj.Current();
 * param_3 := obj.Maximum();
 * param_4 := obj.Minimum();
 */
