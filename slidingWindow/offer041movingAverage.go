package slidingwindow

// https://leetcode.cn/problems/qIsx9U/submissions/

type MovingAverage struct {
	arr []int
	ptr int
	cur int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{arr: make([]int, 0), ptr: size}
}

func (this *MovingAverage) Next(val int) float64 {
	var ans float64
	this.arr = append(this.arr, val)
	l := len(this.arr)
	if l <= this.ptr {
		this.cur += val
		ans = float64(this.cur) / float64(l)
	} else {
		this.cur += val
		this.cur -= this.arr[l-this.ptr-1]
		ans = float64(this.cur) / float64(this.ptr)
	}
	return ans
}
