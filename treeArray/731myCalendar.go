package treearray

// https://leetcode.cn/problems/my-calendar-ii/

type MyCalendarTwo struct {
	oneTime [][2]int
	twoTime [][2]int
}

func ConstructorCal() MyCalendarTwo {
	return MyCalendarTwo{oneTime: make([][2]int, 0), twoTime: make([][2]int, 0)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, v := range this.twoTime {
		if end > v[0] && start < v[1] {
			return false
		}
	}
	for _, v := range this.oneTime {
		if end > v[0] && start < v[1] {
			this.twoTime = append(this.twoTime, [2]int{max(v[0], start), min(v[1], end)})
			if start < v[0] {
				this.oneTime = append(this.oneTime, [2]int{start, v[0]})
			}
			if end > v[1] {
				this.oneTime = append(this.oneTime, [2]int{v[1], end})
			}
		}
	}
	this.oneTime = append(this.oneTime, [2]int{start, end})
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
