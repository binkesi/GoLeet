package priorityqueue

import "container/heap"

// https://leetcode-cn.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) (ans []int) {
	occurrences := map[int]int{}
	for _, num := range nums {
		occurrences[num]++
	}
	var lheap IHeap
	for key, value := range occurrences {
		lheap = append(lheap, [2]int{key, value})
	}
	heap.Init(&lheap)
	for i := 0; i < k; i++ {
		res := heap.Pop(&lheap)
		ans = append(ans, res.([2]int)[0])
	}
	return
}

type IHeap [][2]int

func (h IHeap) Len() int           { return len(h) }
func (h IHeap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h IHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
