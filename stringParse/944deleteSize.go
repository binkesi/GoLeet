package stringparse

// https://leetcode.cn/problems/delete-columns-to-make-sorted/submissions/

func minDeletionSize(strs []string) (ans int) {
	n := len(strs[0])
	for i := 0; i < n; i++ {
		tmp := byte('a')
		for _, s := range strs {
			if s[i] >= byte(tmp) {
				tmp = s[i]
			} else {
				ans++
				break
			}
		}
	}
	return
}
