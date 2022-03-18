package priorityqueue

import "container/heap"

// https://leetcode-cn.com/problems/task-scheduler/

func leastInterval(tasks []byte, n int) (ans int) {
	l := len(tasks)
	if n == 0 {
		return l
	}
	taskMap := make(map[byte]int)
	for _, v := range tasks {
		taskMap[v] += 1
	}
	taskHeap := lheap{}
	for _, v := range taskMap {
		heap.Push(&taskHeap, v)
	}
	for taskHeap.Len() > 0 {
		res := []int{}
		for i := 0; i <= n; i++ {
			if taskHeap.Len() == 0 && len(res) == 0 {
				break
			}
			if taskHeap.Len() > 0 {
				num := heap.Pop(&taskHeap).(int)
				num -= 1
				if num > 0 {
					res = append(res, num)
				}
			}
			ans += 1
		}
		for _, v := range res {
			taskHeap.Push(v)
		}
	}
	return
}

type lheap []int

func (h lheap) Len() int             { return len(h) }
func (h lheap) Less(i1, i2 int) bool { return (h[i1] > h[i2]) }
func (h lheap) Swap(i1, i2 int)      { h[i1], h[i2] = h[i2], h[i1] }
func (h *lheap) Push(v interface{})  { *h = append(*h, v.(int)) }
func (h *lheap) Pop() interface{} {
	old := *h
	n := old.Len()
	res := old[n-1]
	*h = old[:n-1]
	return res
}
