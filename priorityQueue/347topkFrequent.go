package priorityqueue

import "container/heap"

// https://leetcode-cn.com/problems/top-k-frequent-elements/

func topKFrequent(nums []int, k int) (ans []int) {
	nMap := make(map[int]int)
	for _, v := range nums {
		nMap[v] += 1
	}
	numHeap := lowHeap{}
	for k, v := range nMap {
		heap.Push(&numHeap, [2]int{k, v})
	}
	for i := 0; i < k; i++ {
		num := heap.Pop(&numHeap).([2]int)
		ans = append(ans, num[0])
	}
	return
}

type lowHeap [][2]int

func (l lowHeap) Len() int              { return len(l) }
func (l lowHeap) Less(i, j int) bool    { return l[i][1] > l[j][1] }
func (l lowHeap) Swap(i, j int)         { l[i], l[j] = l[j], l[i] }
func (l *lowHeap) Push(num interface{}) { *l = append(*l, num.([2]int)) }
func (l *lowHeap) Pop() interface{} {
	old := *l
	n := old.Len()
	top := old[n-1]
	*l = old[:n-1]
	return top
}
