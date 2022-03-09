package arraymanipulate

// https://leetcode-cn.com/problems/smallest-rotation-with-highest-score/

func bestRotation(nums []int) (ans int) {
	res := make([]int, 100010)
	n := len(nums)
	var add func(int, int)
	add = func(i1, i2 int) {
		res[i1] += 1
		res[i2+1] -= 1
	}
	for i := 0; i < n; i++ {
		a, b := (i-(n-1)+n)%n, (i-nums[i]+n)%n
		if a <= b {
			add(a, b)
		} else {
			add(0, b)
			add(a, n-1)
		}
	}
	tmp := res[0]
	for i := 1; i < 100010; i++ {
		res[i] += res[i-1]
	}
	for i := 1; i < 100010; i++ {
		if res[i] > tmp {
			tmp = res[i]
			ans = i
		}
	}
	return
}
