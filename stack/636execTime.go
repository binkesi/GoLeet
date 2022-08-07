package stack

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	exec := [][]int{}
	ans := make([]int, n)
	for _, s := range logs {
		tmp := strings.Split(s, ":")
		res := []int{}
		id, _ := strconv.Atoi(tmp[0])
		stat := 0
		if tmp[1] == "end" {
			stat = 1
		}
		timeStamp, _ := strconv.Atoi(tmp[2])
		res = append(res, id, stat, timeStamp)
		exec = append(exec, res)
	}
	taskStack := [][]int{exec[0]}
	mn := len(exec)
	for i := 1; i < mn; i++ {
		le := len(taskStack)
		if le == 0 {
			taskStack = append(taskStack, exec[i])
			continue
		}
		period := exec[i][2] - exec[i-1][2]
		id1, id2 := taskStack[le-1][0], exec[i][0]
		ll, l, r := taskStack[le-1][1], exec[i-1][1], exec[i][1]
		if id1 == id2 && ll == 0 && r == 1 {
			if l == 0 {
				ans[id1] += (period + 1)
			} else {
				ans[id1] += period
			}
			taskStack = taskStack[:le-1]
		} else {
			if l == 0 && r == 0 {
				ans[id1] += period
			} else if l == 1 && r == 0 {
				ans[id1] += (period - 1)
			} else {
				ans[id2] += period
			}
			taskStack = append(taskStack, exec[i])
		}
	}
	return ans
}
