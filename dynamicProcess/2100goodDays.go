package dynamicprocess

// https://leetcode-cn.com/problems/find-good-days-to-rob-the-bank/

func goodDaysToRobBank(security []int, time int) (ans []int) {
	n := len(security)
	m := n - time*2
	if m <= 0 {
		return ans
	}
	res := make([][2]int, m)
	for i, k := time-1, 0; i >= 0; i-- {
		if security[i] >= security[i+1] {
			k++
		} else {
			res[0][0] = k
			break
		}
		if i == 0 {
			res[0][0] = k
		}
	}
	for j, k := n-time, 0; j <= n-1; j++ {
		if security[j] >= security[j-1] {
			k++
		} else {
			res[m-1][1] = k
			break
		}
		if j == n-1 {
			res[m-1][1] = k
		}
	}
	for i := time + 1; i <= n-time-1; i++ {
		if security[i] <= security[i-1] {
			res[i-time][0] = res[i-time-1][0] + 1
		} else {
			res[i-time][0] = 0
		}
	}
	for i := n - time - 2; i >= time; i-- {
		if security[i] <= security[i+1] {
			res[i-time][1] = res[i-time+1][1] + 1
		} else {
			res[i-time][1] = 0
		}
	}
	for i, v := range res {
		if v[0] >= time && v[1] >= time {
			ans = append(ans, time+i)
		}
	}
	return
}
