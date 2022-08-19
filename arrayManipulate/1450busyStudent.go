package arraymanipulate

// https://leetcode.cn/problems/number-of-students-doing-homework-at-a-given-time/

func busyStudent(startTime []int, endTime []int, queryTime int) (ans int) {
	l := len(startTime)
	for i := 0; i < l; i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ans++
		}
	}
	return
}

func busyStudent2(startTime []int, endTime []int, queryTime int) (ans int) {
	maxTime := 0
	for _, v := range endTime {
		if v > maxTime {
			maxTime = v
		}
	}
	if queryTime > maxTime {
		return
	}
	cnt := make([]int, maxTime+2)
	for i, v := range startTime {
		cnt[v]++
		cnt[endTime[i]+1]--
	}
	for _, v := range cnt[:queryTime+1] {
		ans += v
	}
	return
}
