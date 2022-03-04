package dfsbfs

// https://leetcode-cn.com/problems/course-schedule-ii/

func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
	posMap := make(map[int]map[int]struct{}, numCourses)
	negMap := make(map[int]map[int]struct{}, numCourses)
	for k := 0; k < numCourses; k++ {
		posMap[k] = make(map[int]struct{})
	}
	for k := 0; k < numCourses; k++ {
		negMap[k] = make(map[int]struct{})
	}
	for _, v := range prerequisites {
		posMap[v[0]][v[1]] = struct{}{}
		negMap[v[1]][v[0]] = struct{}{}
	}
	for len(posMap) != 0 {
		change := false
		for k, v := range posMap {
			if len(v) == 0 {
				change = true
				ans = append(ans, k)
				for key := range negMap[k] {
					delete(posMap[key], k)
				}
				delete(posMap, k)
			}
		}
		if change == false {
			ans = []int{}
			return
		}
	}
	return
}
