package weekgame

// https://leetcode.cn/contest/weekly-contest-300/problems/spiral-matrix-iv/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func spiralMatrix(m int, n int, head *ListNode) [][]int {
	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[i][j] = -1
		}
	}
	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	ind := 0
	for i, j := 0, 0; ; {
		val := head.Val
		ans[i][j] = val
		if head.Next != nil {
			head = head.Next
		} else {
			break
		}
		xi, xj := i+dirs[ind][0], j+dirs[ind][1]
		if xi == m || xi < 0 || xj == n || xj < 0 || ans[xi][xj] != -1 {
			ind = (ind + 1) % 4
			xi, xj = i+dirs[ind][0], j+dirs[ind][1]
		}
		i, j = xi, xj
	}
	return ans
}
