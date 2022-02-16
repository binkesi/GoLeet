package priorityqueue

import "container/heap"

// https://leetcode-cn.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) (ans []int) {
	nMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := nMap[v]; !ok {
			nMap[v] = 1
		} else {
			nMap[v] += 1
		}
	}
	nlis := make([][2]int, 0)
	for k, v := range nMap {
		nlis = append(nlis, [2]int{v, k})
	}
	ll := lheap(nlis)
	heap.Init(&ll)
	for i := 1; i <= k; i++ {
		ans = append(ans, ll.Pop().([2]int)[1])
		heap.Fix(ll, 0)
	}
	return
}

type lheap [][2]int

func (h lheap) Len() int { return len(h) }
func (h lheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h lheap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h *lheap) Push(a interface{}) { *h = append(*h, a.([2]int)) }
func (h *lheap) Pop() interface{} {
	old := *h
	res := old[0]
	*h = old[1:]
	return res
}
