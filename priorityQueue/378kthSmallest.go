package priorityqueue

import "container/heap"

func kthSmallest(matrix [][]int, k int) (ans int) {
	n := len(matrix)
	numHeap := heapS{}
	for x := 0; x < n; x++ {
		heap.Push(&numHeap, numInfo{x, 0, matrix[x][0]})
	}
	for i := 0; i < k; i++ {
		res := heap.Pop(&numHeap).(numInfo)
		if i == k-1 {
			ans = res.val
			break
		}
		if res.y != n-1 {
			heap.Push(&numHeap, numInfo{res.x, res.y + 1, matrix[res.x][res.y+1]})
		}
	}
	return
}

type numInfo struct {
	x   int
	y   int
	val int
}

type heapS []numInfo

func (h heapS) Len() int              { return len(h) }
func (h heapS) Less(i, j int) bool    { return h[i].val < h[j].val }
func (h heapS) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *heapS) Push(num interface{}) { *h = append(*h, num.(numInfo)) }
func (h *heapS) Pop() interface{} {
	old := *h
	n := h.Len()
	top := old[n-1]
	*h = old[:n-1]
	return top
}
