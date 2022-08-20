package classdesign

// https://leetcode.cn/problems/design-an-ordered-stream/

type OrderedStream struct {
	ptr  int
	sArr []string
}

func ConstructorOrdered(n int) OrderedStream {
	return OrderedStream{ptr: 1, sArr: make([]string, n+1)}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	ans := []string{}
	this.sArr[idKey] = value
	l := len(this.sArr)
	for this.ptr < l && this.sArr[this.ptr] != "" {
		ans = append(ans, this.sArr[this.ptr])
		this.ptr++
	}
	return ans
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(idKey,value);
 */
