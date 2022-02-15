package weekgame

import (
	"container/heap"
	"sort"
)

// https://leetcode-cn.com/problems/minimum-difference-in-sums-after-removal-of-elements/

func minimumDifference(nums []int) (ans int64) {
	lenN := len(nums)
	n := lenN / 3
	minh := minHeap{nums[2*n:]}
	curminSum := 0
	for i := 2 * n; i < lenN; i++ {
		curminSum += nums[i]
	}
	heap.Init(&minh)
	maxh := maxHeap{nums[0:n]}
	curmaxSum := 0
	for i := 0; i < n; i++ {
		curmaxSum += nums[i]
	}
	heap.Init(&maxh)
	minSum := []int{curminSum}
	for i := 2*n - 1; i >= n; i-- {
		curMin := minh.IntSlice[0]
		if nums[i] > curMin {
			minh.IntSlice[0] = nums[i]
			heap.Fix(minh, 0)
			curminSum = curminSum - curMin + nums[i]
		}
		minSum = append(minSum, curminSum)
	}
	ans = int64(curmaxSum - minSum[n-1])

	for i := n; i <= 2*n; i++ {
		curDif := curmaxSum - minSum[2*n-i]
		if int64(curDif) < ans {
			ans = int64(curDif)
		}
		curMax := maxh.IntSlice[0]
		if nums[i] < curMax {
			maxh.IntSlice[0] = nums[i]
			heap.Fix(maxh, 0)
			curmaxSum = curmaxSum - curMax + nums[i]
		}
	}
	return
}

type minHeap struct {
	sort.IntSlice
}

func (minHeap) Push(interface{})     {}
func (minHeap) Pop() (_ interface{}) { return }

type maxHeap struct {
	sort.IntSlice
}

func (m maxHeap) Less(i, j int) bool { return m.IntSlice[i] > m.IntSlice[j] }
func (maxHeap) Push(interface{})     {}
func (maxHeap) Pop() (_ interface{}) { return }
