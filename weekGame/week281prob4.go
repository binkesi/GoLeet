package weekgame

// https://leetcode-cn.com/contest/weekly-contest-281/problems/count-array-pairs-divisible-by-k/

func coutPairs(nums []int, k int) (ans int64) {
	res := make(map[int]map[int]struct{})
	for div := k / 2; div >= 1; div-- {
		d1, d2 := 0, 0
		n1, n2 := make(map[int]int), make(map[int]int)
		if k%div == 0 {
			d1 = div
			d2 = k / div
		}
		for i, v := range nums {
			if v%d1 == 0 {
				n1[i] = v
			}
			if v%d2 == 0 {
				n2[i] = v
			}
		}
		for k1, _ := range n1 {
			for k2, _ := range n2 {
				if k1 < k2 {
					res[k1][k2] = struct{}{}
				} else {
					res[k2][k1] = struct{}{}
				}
			}
		}
	}
	for _, vout := range res {
		ans += int64(len(vout))
	}
	return
}
