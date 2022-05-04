package arraymanipulate

// https://leetcode.cn/problems/find-the-winner-of-the-circular-game/submissions/

func findTheWinner(n int, k int) int {
	parts := make([]int, n+1)
	for i := range parts {
		parts[i] = i
	}
	res := parts[1:]
	for i, cur := 1, 0; len(res) > 1; i++ {
		l := len(res)
		del := (cur + k - 1) % l
		if del != l-1 {
			res = append(res[0:del], res[del+1:]...)
		} else {
			res = res[0:del]
		}
		cur = del % len(res)
	}
	return res[0]
}
