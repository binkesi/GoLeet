package hashmap

// https://leetcode.cn/problems/group-the-people-given-the-group-size-they-belong-to/

func groupThePeople(groupSizes []int) (ans [][]int) {
	groupMap := make(map[int][]int)
	for i := range groupSizes {
		size := groupSizes[i]
		if _, ok := groupMap[size]; !ok {
			groupMap[size] = []int{len(ans)}
			ans = append(ans, []int{i})
		} else {
			for k, v := range groupMap[size] {
				if len(ans[v]) < size {
					ans[v] = append(ans[v], i)
					break
				}
				if k == len(groupMap[size])-1 {
					groupMap[size] = append(groupMap[size], len(ans))
					ans = append(ans, []int{i})
				}
			}
		}
	}
	return
}
