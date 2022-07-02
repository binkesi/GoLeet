package priorityqueue

import "container/heap"

// https://leetcode.cn/problems/minimum-number-of-refueling-stops/submissions/

func minRefuelStops(target int, startFuel int, stations [][]int) (ans int) {
	loc, index := startFuel, 0
	n := len(stations)
	myHeap := lheap{}
	for loc < target {
		for index < n && stations[index][0] <= loc {
			heap.Push(&myHeap, stations[index][1])
			index++
		}
		if myHeap.Len() == 0 {
			return -1
		}
		maxFuel := heap.Pop(&myHeap).(int)
		loc += maxFuel
		ans++
	}
	return
}

type lheap []int

func (h lheap) Len() int            { return len(h) }
func (h lheap) Less(i, j int) bool  { return h[i] > h[j] }
func (h lheap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *lheap) Push(i interface{}) { *h = append(*h, i.(int)) }
func (h *lheap) Pop() interface{} {
	old := *h
	n := old.Len()
	res := old[n-1]
	*h = old[0 : n-1]
	return res
}
