package stringparse

// https://leetcode-cn.com/problems/longest-absolute-file-path/solution/

func lengthLongestPath(input string) (ans int) {
	n := len(input)
	level := make([]int, n+1)
	for i := 0; i < n; {
		// 检测当前文件的深度
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++
		}

		// 统计当前文件名的长度
		length, isFile := 0, false
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		i++ // 跳过换行符

		if depth > 1 {
			length += level[depth-1] + 1
		}
		if isFile {
			ans = max(ans, length)
		} else {
			level[depth] = length
		}
	}
	return
}
