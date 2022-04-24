package twopointer

// https://leetcode-cn.com/problems/binary-gap/submissions/

func binaryGap(n int) (ans int) {
	pre, cur := -1, 0
	for n > 0 {
		if n&1 == 1 && pre != -1 {
			dis := cur - pre
			pre = cur
			if dis > ans {
				ans = dis
			}
		} else if n&1 == 1 && pre == -1 {
			pre = cur
		}
		n >>= 1
		cur++
	}
	return
}
