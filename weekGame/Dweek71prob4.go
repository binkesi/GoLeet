package weekgame

import "container/heap"

// https://leetcode-cn.com/problems/minimum-difference-in-sums-after-removal-of-elements/

func minimumDifference(nums []int) (ans int64) {
	lenN := len(nums)
	n := lenN / 3
	numsF, numsM, numsS := nums[0:n], nums[n:2*n], nums[2*n:]
	for i := 0; i < n; i++ {
		ans = ans + int64(numsF[i]-numsS[i])
	}
	hl, hs := &lHeap{}, &sHeap{}
	for i := 0; i < n; i++ {
		heap.Push(hl, numsF[i])
		heap.Push(hs, numsS[i])
	}
	for j, l, r := 0, 0, n-1; j < n; j++ {
		hlnum, hsnum := heap.Pop(hl), heap.Pop(hs)
		lnum, rnum := numsM[l], numsM[r]
		if hlnum.(int)-lnum >= rnum-hsnum.(int) {
			heap.Push(hl, lnum)
			heap.Push(hs, hsnum)
			l += 1
			ans -= int64(hlnum.(int) - lnum)
		} else {
			heap.Push(hs, rnum)
			heap.Push(hl, hlnum)
			r -= 1
			ans -= int64(rnum - hsnum.(int))
		}
	}
	return
}

type lHeap []int

func (h lHeap) Len() int { return len(h) }
func (h lHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h lHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *lHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *lHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

type sHeap []int

func (h sHeap) Len() int { return len(h) }
func (h sHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h sHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *sHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *sHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
