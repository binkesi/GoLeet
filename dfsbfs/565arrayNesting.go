package dfsbfs

// https://leetcode.cn/problems/array-nesting/solution/

func arrayNesting(nums []int) (ans int) {
	l := len(nums)
	visited := make([]int, l)
	var dfs func(start, ind int, cnt *int)
	dfs = func(start, ind int, cnt *int) {
		visited[ind] = 1
		*cnt += 1
		if nums[ind] == start {
			if *cnt > ans {
				ans = *cnt
			}
			return
		}
		dfs(start, nums[ind], cnt)
	}
	for i := 0; i < l; i++ {
		if visited[i] == 1 {
			continue
		}
		cnt := 0
		dfs(i, i, &cnt)
	}
	return
}
