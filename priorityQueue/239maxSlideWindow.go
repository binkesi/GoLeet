package priorityqueue

import "container/heap"

// https://leetcode-cn.com/problems/sliding-window-maximum/solution/

func maxSlidingWindow(nums []int, k int) (ans []int) {
	l := len(nums)
	lheap := heapL{}
	heap.Init(&lheap)
	for i := 0; i < k; i++ {
		heap.Push(&lheap, [2]int{i, nums[i]})
	}
	val := lheap[0]
	ans = append(ans, val[1])
	for j := 1; j <= l-k; j++ {
		heap.Push(&lheap, [2]int{j + k - 1, nums[j+k-1]})
		maxV := lheap[0]
		for maxV[0] < j {
			_ = heap.Pop(&lheap).([2]int)
			maxV = lheap[0]
		}
		ans = append(ans, maxV[1])
	}
	return
}

type heapL [][2]int

func (h heapL) Len() int            { return len(h) }
func (h heapL) Less(i, j int) bool  { return h[i][1] > h[j][1] }
func (h heapL) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *heapL) Push(v interface{}) { *h = append(*h, v.([2]int)) }
func (h *heapL) Pop() interface{} {
	old := *h
	n := old.Len()
	res := old[n-1]
	*h = old[0 : n-1]
	return res
}
