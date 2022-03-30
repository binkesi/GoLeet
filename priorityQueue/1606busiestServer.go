package priorityqueue

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://leetcode-cn.com/problems/find-servers-that-handled-most-number-of-requests/

func busiestServers(k int, arrival []int, load []int) (ans []int) {
	cnt := make([][2]int, k)
	for i := range cnt {
		cnt[i] = [2]int{i, 0}
	}
	avaliableServer := make([]int, k)
	n := len(arrival)
	for i := range avaliableServer {
		avaliableServer[i] = i
	}
	var runningServer serHeap
	var findServer func(int)
	var releaseServer func(int)
	findServer = func(index int) {
		fNum := index % k
		servNum := avaliableServer[fNum]
		if servNum != -1 {
			releaseTime := arrival[index] + load[index]
			heap.Push(&runningServer, [2]int{servNum, releaseTime})
			nextAva := avaliableServer[(fNum+1)%k]
			if nextAva == servNum {
				for i := range avaliableServer {
					avaliableServer[i] = -1
				}
			} else {
				for i := fNum; avaliableServer[i] == fNum; {
					avaliableServer[i] = nextAva
					i = (i + k - 1) % k
				}
			}
			cnt[servNum][1] += 1
		} else {
			return
		}
	}
	releaseServer = func(t int) {
		if runningServer.Len() == 0 || runningServer[0][1] > t {
			return
		} else {
			serv := heap.Pop(&runningServer).([2]int)[0]
			if avaliableServer[serv] == -1 {
				for i := range avaliableServer {
					avaliableServer[i] = serv
				}
			} else {
				avaliableServer[serv] = serv
				for i := serv; avaliableServer[i] > serv; {
					avaliableServer[i] = serv
					i = (i + k - 1) % k
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		if i >= 1 {
			for t := arrival[i-1]; t <= arrival[i]; t++ {
				releaseServer(t)
			}
			fmt.Println("after release:", avaliableServer)
		}
		findServer(i)
		fmt.Println(avaliableServer)
	}
	fmt.Println(cnt)
	sort.Slice(cnt, func(i, j int) bool { return cnt[i][1] < cnt[j][1] })
	maxV := cnt[k-1][1]
	ans = append(ans, cnt[k-1][0])
	for i := k - 2; i >= 0; i-- {
		if cnt[i][1] == maxV {
			ans = append(ans, cnt[i][0])
		} else {
			break
		}
	}
	return
}

type serHeap [][2]int

func (h serHeap) Len() int               { return len(h) }
func (h serHeap) Less(i, j int) bool     { return h[i][1] < h[j][1] }
func (h serHeap) Swap(i, j int)          { h[i], h[j] = h[j], h[i] }
func (h *serHeap) Push(serv interface{}) { *h = append(*h, serv.([2]int)) }
func (h *serHeap) Pop() interface{} {
	old := *h
	n := old.Len()
	serv := old[n-1]
	*h = old[:n-1]
	return serv
}
