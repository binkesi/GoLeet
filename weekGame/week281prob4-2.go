package weekgame

// https://leetcode-cn.com/contest/weekly-contest-281/problems/count-array-pairs-divisible-by-k/

func coutPairs2(nums []int, k int) (ans int64) {
	gcdArr := make([]int, k+1)
	for _, v := range nums {
		gcdArr[gcd(v, k)] += 1
	}
	for i, v := range gcdArr {
		if i > k/2 {
			break
		}
		if v > 0 {
			for j := 1; j <= i; j++ {
				if j*k/i == i {
					ans += int64(v * (v - 1) / 2) // 同一个因子的时候要注意会重复算相同下标
				} else if j*k/i > i {
					ans += int64(v * gcdArr[j*k/i]) // 不同因子只算比第一个因子大的
				}
			}
		}
	}
	ans += int64(gcdArr[k] * (gcdArr[k] - 1) / 2) // 注意再加上因子k的
	return
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
