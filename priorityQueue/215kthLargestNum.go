package priorityqueue

import (
	"container/heap"
	"sort"
)

// https://leetcode-cn.com/problems/kth-largest-element-in-an-array/

func findKthLargest(nums []int, k int) int {
	lenN := len(nums)
	minh := minHeap{nums[0:k]}
	heap.Init(minh)
	for i := k; i < lenN; i++ {
		if nums[i] > minh.IntSlice[0] {
			minh.IntSlice[0] = nums[i]
			heap.Fix(minh, 0)
		}
	}
	return minh.IntSlice[0]
}

type minHeap struct {
	sort.IntSlice
}

func (minHeap) Push(interface{})     {}
func (minHeap) Pop() (_ interface{}) { return }
