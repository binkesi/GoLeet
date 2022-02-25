package dfsbfs

// https://leetcode-cn.com/problems/course-schedule/

func canFinish(numCourses int, prerequisites [][]int) bool {
	courseMap := make(map[int]map[int]struct{}, numCourses)
	for i := 0; i < numCourses; i++ {
		courseMap[i] = make(map[int]struct{})
	}
	for _, v := range prerequisites {
		courseMap[v[0]][v[1]] = struct{}{}
	}
	pop := -1
	for len(courseMap) != 0 {
		for k, v := range courseMap {
			if len(v) == 0 {
				pop = k
				delete(courseMap, k)
				break
			}
		}
		if pop == -1 {
			return false
		}
		for _, v := range courseMap {
			if _, ok := v[pop]; ok {
				delete(v, pop)
			}
		}
		pop = -1
	}
	return true
}
