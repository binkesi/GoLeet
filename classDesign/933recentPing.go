package classdesign

// https://leetcode.cn/problems/number-of-recent-calls/submissions/

type RecentCounter struct {
	counter []int
}

func Constructor3() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	(*this).counter = append((*this).counter, t)
	for (*this).counter[0] < t-3000 {
		(*this).counter = (*this).counter[1:]
	}
	return len((*this).counter)
}
